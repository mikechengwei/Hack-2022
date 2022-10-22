package oracle

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/knullhhf/hack22/pkg/net/storage"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pingcap/log"
	"github.com/shopspring/decimal"
	"github.com/thinkeridea/go-extend/exstrings"
	"github.com/xwb1989/sqlparser"
	"go.uber.org/zap"
)

// Dumper is the dump progress structure
type Dumper struct {
	conf *Config

	ctx      context.Context
	extStore *storage.SocketStorage
	dbHandle *sql.DB
}

func NewDumper(ctx context.Context, conf *Config) *Dumper {
	oradb, err := NewOracleDBEngine(conf)
	if err != nil {
		panic(err)
	}
	return &Dumper{
		conf:     conf,
		ctx:      ctx,
		extStore: conf.ExtStorage,
		dbHandle: oradb,
	}

}

func (d *Dumper) Dump() error {
	stmt, err := sqlparser.Parse(d.conf.SQL)
	if err != nil {
		return err
	}
	sel := stmt.(*sqlparser.Select)
	sel.From = sqlparser.TableExprs{
		sel.From[0].(*sqlparser.AliasedTableExpr).RemoveHints(),
	}
	buf := sqlparser.NewTrackedBuffer(nil)
	sel.From.Format(buf)

	dbTableArr := strings.Split(buf.String(), ".")

	buf = sqlparser.NewTrackedBuffer(nil)
	sel.SelectExprs.Format(buf)
	selectExprs := buf.String()

	// select ID,NAME from marvin.table1
	if strings.EqualFold(selectExprs, "*") {
		selectExpr, err := d.AdjustTableSelectColumn(dbTableArr[0], dbTableArr[1])
		if err != nil {
			return err
		}
		d.conf.SQL = strings.Replace(d.conf.SQL, "*", selectExpr, 1)
	}

	if err := d.ExtractorTableRecord(d.conf.SQL); err != nil {
		return err
	}
	return nil
}

func (d *Dumper) AdjustTableSelectColumn(schemaName, tableName string) (string, error) {
	columnInfo, err := d.GetOracleTableColumn(schemaName, tableName)
	if err != nil {
		return "", err
	}

	var columnNames []string

	for _, rowCol := range columnInfo {
		switch strings.ToUpper(rowCol["DATA_TYPE"]) {
		// 数字
		case "NUMBER":
			columnNames = append(columnNames, rowCol["COLUMN_NAME"])
		case "DECIMAL", "DEC", "DOUBLE PRECISION", "FLOAT", "INTEGER", "INT", "REAL", "NUMERIC", "BINARY_FLOAT", "BINARY_DOUBLE", "SMALLINT":
			columnNames = append(columnNames, rowCol["COLUMN_NAME"])
		// 字符
		case "BFILE", "CHARACTER", "LONG", "NCHAR VARYING", "ROWID", "UROWID", "VARCHAR", "XMLTYPE", "CHAR", "NCHAR", "NVARCHAR2", "NCLOB", "CLOB":
			columnNames = append(columnNames, rowCol["COLUMN_NAME"])
		// 二进制
		case "BLOB", "LONG RAW", "RAW":
			columnNames = append(columnNames, rowCol["COLUMN_NAME"])
		// 时间
		case "DATE":
			columnNames = append(columnNames, StringsBuilder("TO_CHAR(", rowCol["COLUMN_NAME"], ",'yyyy-MM-dd HH24:mi:ss') AS ", rowCol["COLUMN_NAME"]))
		// 默认其他类型
		default:
			if strings.Contains(rowCol["DATA_TYPE"], "INTERVAL") {
				columnNames = append(columnNames, StringsBuilder("TO_CHAR(", rowCol["COLUMN_NAME"], ") AS ", rowCol["COLUMN_NAME"]))
			} else if strings.Contains(rowCol["DATA_TYPE"], "TIMESTAMP") {
				dataScale, err := strconv.Atoi(rowCol["DATA_SCALE"])
				if err != nil {
					return "", fmt.Errorf("aujust oracle timestamp datatype scale [%s] strconv.Atoi failed: %v", rowCol["DATA_SCALE"], err)
				}
				if dataScale == 0 {
					columnNames = append(columnNames, StringsBuilder("TO_CHAR(", rowCol["COLUMN_NAME"], ",'yyyy-mm-dd hh24:mi:ss') AS ", rowCol["COLUMN_NAME"]))
				} else if dataScale < 0 && dataScale <= 6 {
					columnNames = append(columnNames, StringsBuilder("TO_CHAR(", rowCol["COLUMN_NAME"],
						",'yyyy-mm-dd hh24:mi:ss.ff", rowCol["DATA_SCALE"], "') AS ", rowCol["COLUMN_NAME"]))
				} else {
					columnNames = append(columnNames, StringsBuilder("TO_CHAR(", rowCol["COLUMN_NAME"], ",'yyyy-mm-dd hh24:mi:ss.ff6') AS ", rowCol["COLUMN_NAME"]))
				}

			} else {
				columnNames = append(columnNames, rowCol["COLUMN_NAME"])
			}
		}

	}

	return strings.Join(columnNames, ","), nil
}

func (d *Dumper) ExtractorTableRecord(oracleQuery string) error {
	startTime := time.Now()

	rows, err := d.dbHandle.Query(oracleQuery)
	if err != nil {
		return fmt.Errorf("get oracle schema table record by sql falied: %v", err)
	}

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("[%v] error on general query rows.Columns failed", err.Error())
	}

	if d.conf.Header {
		// id,name
		if _, err := d.extStore.Writer.Write(d.ctx, []byte(StringsBuilder(exstrings.Join(columns, d.conf.Separator), d.conf.Terminator))); err != nil {
			return fmt.Errorf("failed to write headers: %v", err)
		}
	}

	var columnTypes []string
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return fmt.Errorf("failed to csv get rows columnTypes: %v", err)
	}

	for _, ct := range colTypes {
		// 数据库字段类型 DatabaseTypeName() 映射 go 类型 ScanType()
		columnTypes = append(columnTypes, ct.ScanType().String())
	}

	// 数据 SCAN
	rawResult := make([][]byte, len(columns))
	dest := make([]interface{}, len(columns))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	// 表行数读取
	for rows.Next() {

		var results []string

		err = rows.Scan(dest...)
		if err != nil {
			return err
		}

		for i, raw := range rawResult {
			// 注意 Oracle/Mysql NULL VS 空字符串区别
			// Oracle 空字符串与 NULL 归于一类，统一 NULL 处理 （is null 可以查询 NULL 以及空字符串值，空字符串查询无法查询到空字符串值）
			// Mysql 空字符串与 NULL 非一类，NULL 是 NULL，空字符串是空字符串（is null 只查询 NULL 值，空字符串查询只查询到空字符串值）
			// 按照 Oracle 特性来，转换同步统一转换成 NULL 即可，但需要注意业务逻辑中空字符串得写入，需要变更
			// Oracle/Mysql 对于 'NULL' 统一字符 NULL 处理，查询出来转成 NULL,所以需要判断处理
			if raw == nil {
				results = append(results, "NULL")
			} else if string(raw) == "" {
				results = append(results, "NULL")
			} else {
				switch columnTypes[i] {
				case "int64":
					r, err := StrconvIntBitSize(string(raw), 64)
					if err != nil {
						return err
					}
					results = append(results, fmt.Sprintf("%v", r))
				case "uint64":
					r, err := StrconvUintBitSize(string(raw), 64)
					if err != nil {
						return err
					}
					results = append(results, fmt.Sprintf("%v", r))
				case "float32":
					r, err := StrconvFloatBitSize(string(raw), 32)
					if err != nil {
						return err
					}
					results = append(results, fmt.Sprintf("%v", r))
				case "float64":
					r, err := StrconvFloatBitSize(string(raw), 64)
					if err != nil {
						return err
					}
					results = append(results, fmt.Sprintf("%v", r))
				case "rune":
					r, err := StrconvRune(string(raw))
					if err != nil {
						return err
					}
					results = append(results, fmt.Sprintf("%v", r))
				case "godror.Number":
					r, err := decimal.NewFromString(string(raw))
					if err != nil {
						return err
					}
					if r.IsInteger() {
						si, err := StrconvIntBitSize(string(raw), 64)
						if err != nil {
							return err
						}
						results = append(results, fmt.Sprintf("%v", si))
					} else {
						rf, err := StrconvFloatBitSize(string(raw), 64)
						if err != nil {
							return err
						}
						results = append(results, fmt.Sprintf("%v", rf))
					}
				default:
					var (
						by []byte
						bs string
					)
					// 处理字符集、特殊字符转义、字符串引用定界符
					if strings.ToUpper(d.conf.Charset) == "GBK" {
						gbkBytes, err := Utf8ToGbk(raw)
						if err != nil {
							return err
						}
						by = gbkBytes
					} else {
						by = raw
					}

					if d.conf.EscapeBackslash {
						bs = SpecialLetters(by)
					} else {
						bs = string(by)
					}

					if d.conf.Delimiter == "" {
						results = append(results, bs)
					} else {
						results = append(results, StringsBuilder(d.conf.Delimiter, bs, d.conf.Delimiter))
					}
				}
			}
		}

		// 数据写入
		if _, err = d.extStore.Writer.Write(d.ctx, []byte(StringsBuilder(exstrings.Join(results, d.conf.Separator), d.conf.Terminator))); err != nil {
			return fmt.Errorf("failed to write data row to csv %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	// Close Rows
	if err := rows.Close(); err != nil {
		return err
	}

	endTime := time.Now()
	log.L().Info("single full table data extractor finished",
		zap.String("sql", oracleQuery),
		zap.String("cost", endTime.Sub(startTime).String()))

	return nil
}

func (d *Dumper) GetOracleTableColumn(schemaName string, tableName string) ([]map[string]string, error) {
	var querySQL string

	querySQL = fmt.Sprintf(`select t.COLUMN_NAME,
	    t.DATA_TYPE,
		 t.CHAR_LENGTH,
		 NVL(t.CHAR_USED,'UNKNOWN') CHAR_USED,
	    NVL(t.DATA_LENGTH,0) AS DATA_LENGTH,
	    DECODE(NVL(TO_CHAR(t.DATA_PRECISION),'*'),'*','38',TO_CHAR(t.DATA_PRECISION)) AS DATA_PRECISION,
	    DECODE(NVL(TO_CHAR(t.DATA_SCALE),'*'),'*','127',TO_CHAR(t.DATA_SCALE)) AS DATA_SCALE,
		t.NULLABLE,
	    t.DATA_DEFAULT,
	    c.COMMENTS
	from dba_tab_columns t, dba_col_comments c
	where t.table_name = c.table_name
	and t.column_name = c.column_name
	and t.owner = c.owner
	and upper(t.owner) = upper('%s')
	and upper(t.table_name) = upper('%s')
	order by t.COLUMN_ID`,
		strings.ToUpper(schemaName),
		strings.ToUpper(tableName))

	_, queryRes, err := Query(d.dbHandle, querySQL)
	if err != nil {
		return queryRes, err
	}
	if len(queryRes) == 0 {
		return queryRes, fmt.Errorf("oracle table [%s.%s] column info cann't be null", schemaName, tableName)
	}

	// check constraints notnull
	// search_condition long datatype
	_, condRes, err := Query(d.dbHandle, fmt.Sprintf(`SELECT
				col.COLUMN_NAME,
				cons.SEARCH_CONDITION
				FROM
				DBA_CONS_COLUMNS col,
				DBA_CONSTRAINTS cons
				WHERE
				col.OWNER = cons.OWNER
				AND col.TABLE_NAME = cons.TABLE_NAME
				AND col.CONSTRAINT_NAME = cons.CONSTRAINT_NAME
				AND cons.CONSTRAINT_TYPE = 'C'
				AND upper(col.OWNER) = '%s'
				AND upper(col.TABLE_NAME) = '%s'`, strings.ToUpper(schemaName), strings.ToUpper(tableName)))
	if err != nil {
		return queryRes, err
	}

	if len(condRes) == 0 {
		return queryRes, nil
	}

	rep, err := regexp.Compile(`(^.*)(?i:IS NOT NULL)`)
	if err != nil {
		return queryRes, fmt.Errorf("check notnull constraint regexp complile failed: %v", err)
	}
	for _, r := range queryRes {
		for _, c := range condRes {
			if r["COLUMN_NAME"] == c["COLUMN_NAME"] && r["NULLABLE"] == "Y" {
				// 检查约束非空检查
				if rep.MatchString(c["SEARCH_CONDITION"]) {
					r["NULLABLE"] = "N"
				}
			}
		}
	}
	return queryRes, nil
}
