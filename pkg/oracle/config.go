package oracle

import (
	"database/sql"
	"fmt"
	"github.com/knullhhf/hack22/pkg/net/storage"
	"os"
	"runtime"

	"github.com/godror/godror"
	"github.com/godror/godror/dsn"
)

type Config struct {
	Username        string                 `toml:"username" json:"username"`
	Password        string                 `toml:"password" json:"password"`
	Host            string                 `toml:"host" json:"host"`
	Port            int                    `toml:"port" json:"port"`
	ServiceName     string                 `toml:"service-name" json:"service-name"`
	LibDir          string                 `toml:"lib-dir" json:"lib-dir"`
	ConnectParams   string                 `toml:"connect-params" json:"connect-params"`
	SessionParams   []string               `toml:"session-params" json:"session-params"`
	Header          bool                   `toml:"header" json:"header"`
	Separator       string                 `toml:"separator" json:"separator"`
	Terminator      string                 `toml:"terminator" json:"terminator"`
	Delimiter       string                 `toml:"delimiter" json:"delimiter"`
	EscapeBackslash bool                   `toml:"escape-backslash" json:"escape-backslash"`
	Charset         string                 `toml:"charset" json:"charset"`
	SQL             string                 `toml:"sql" json:"sql"`
	ExtStorage      *storage.SocketStorage `json:"-"`
}

// DefaultConfig returns the default export Config for dumpling
func DefaultConfig() *Config {
	return &Config{
		Username:        "marvin",
		Password:        "marvin",
		Host:            "127.0.0.1",
		Port:            1521,
		ServiceName:     "orcl",
		LibDir:          "/Users/marvin/storehouse/oracle/instantclient_19_8", // oracle instance client dir
		ConnectParams:   "",
		SessionParams:   []string{},
		Header:          true,
		Separator:       ",",
		Terminator:      "\n",
		Delimiter:       "\"",
		EscapeBackslash: true,
		Charset:         "utf8",
		SQL:             "",
		ExtStorage:      nil,
	}
}

func NewOracleDBEngine(cfg *Config) (*sql.DB, error) {
	var (
		connString string
		oraDSN     dsn.ConnectionParams
		err        error
	)

	connString = fmt.Sprintf("oracle://%s:%s@%s/%s?connectionClass=POOL_CONNECTION_CLASS&heterogeneousPool=1&%s",
		cfg.Username, cfg.Password, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), cfg.ServiceName, cfg.ConnectParams)
	oraDSN, err = godror.ParseDSN(connString)
	if err != nil {
		return nil, err
	}

	oraDSN.OnInitStmts = cfg.SessionParams

	// libDir won't have any effect on Linux for linking reasons to do with Oracle's libnnz library that are proving to be intractable.
	// You must set LD_LIBRARY_PATH or run ldconfig before your process starts.
	// This is documented in various places for other drivers that use ODPI-C. The parameter works on macOS and Windows.
	switch runtime.GOOS {
	case "linux":
		if err = os.Setenv("LD_LIBRARY_PATH", cfg.LibDir); err != nil {
			return nil, fmt.Errorf("set LD_LIBRARY_PATH env failed: %v", err)
		}
	case "windows", "darwin":
		oraDSN.LibDir = cfg.LibDir
	}

	// godror logger 日志输出
	// godror.SetLogger(zapr.NewLogger(zap.L()))

	sqlDB := sql.OpenDB(godror.NewConnector(oraDSN))
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(0)
	sqlDB.SetConnMaxLifetime(0)

	err = sqlDB.Ping()
	if err != nil {
		return sqlDB, fmt.Errorf("error on ping oracle database connection:%v", err)
	}
	return sqlDB, nil
}
