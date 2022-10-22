package oracle

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Query(db *sql.DB, querySQL string) ([]string, []map[string]string, error) {
	var (
		cols []string
		res  []map[string]string
	)
	rows, err := db.Query(querySQL)
	if err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query failed: [%v]", querySQL, err.Error())
	}
	defer rows.Close()

	//不确定字段通用查询，自动获取字段名称
	cols, err = rows.Columns()
	if err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query rows.Columns failed: [%v]", querySQL, err.Error())
	}

	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return cols, res, fmt.Errorf("general sql [%v] query rows.Scan failed: [%v]", querySQL, err.Error())
		}

		row := make(map[string]string)
		for k, v := range values {
			// 数据库类型 MySQL NULL 是 NULL，空字符串是空字符串
			// 数据库类型 Oracle NULL、空字符串归于一类 NULL
			// Oracle/Mysql 对于 'NULL' 统一字符 NULL 处理，查询出来转成 NULL,所以需要判断处理
			if v == nil { // 处理 NULL 情况，当数据库类型 MySQL 等于 nil
				row[cols[k]] = "NULLABLE"
			} else {
				// 处理空字符串以及其他值情况
				// 数据统一 string 格式显示
				row[cols[k]] = string(v)
			}
		}
		res = append(res, row)
	}

	if err = rows.Err(); err != nil {
		return cols, res, fmt.Errorf("general sql [%v] query rows.Next failed: [%v]", querySQL, err.Error())
	}
	return cols, res, nil
}

func SpecialLetters(bs []byte) string {

	var (
		b     strings.Builder
		chars []rune
	)
	for _, r := range bytes.Runes(bs) {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsSpace(r) {
			// mysql/tidb % 字符, /% 代表 /%，% 代表 % ,无需转义
			if r == '%' {
				chars = append(chars, r)
			} else {
				chars = append(chars, '\\', r)
			}
		} else {
			chars = append(chars, r)
		}
	}

	b.WriteString(string(chars))

	return b.String()
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func StringsBuilder(str ...string) string {
	var b strings.Builder
	for _, p := range str {
		b.WriteString(p)
	}
	return b.String() // no copying
}

func StrconvIntBitSize(s string, bitSize int) (int64, error) {
	i, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return i, err
	}
	return i, nil
}

func StrconvUintBitSize(s string, bitSize int) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return i, err
	}
	return i, nil
}

func StrconvFloatBitSize(s string, bitSize int) (float64, error) {
	i, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return i, err
	}
	return i, nil
}

func StrconvRune(s string) (int32, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return rune(r), err
	}
	return rune(r), nil
}
