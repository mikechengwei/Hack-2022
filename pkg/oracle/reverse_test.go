package oracle

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestRunning(t *testing.T) {
	mysqlConf := DefaultMySQLConfig()
	mysqlConf.Host = "10.2.103.30"
	mysqlConf.Port = 4000
	mysqlConf.Username = "root"
	mysqlConf.Password = ""
	mysqlConf.SchemaName = "marvin"

	oraConf := DefaultOracleConfig()
	oraConf.Username = "marvin"
	oraConf.Password = "marvin"
	oraConf.Host = "10.2.103.31"
	oraConf.Port = 1521
	oraConf.ServiceName = "orclpdb1"
	oraConf.LibDir = "/Users/marvin/storehouse/oracle/instantclient_19_8"
	oraConf.ConnectParams = "poolMinSessions=50&poolMaxSessions=1000&poolWaitTimeout=360s&poolSessionMaxLifetime=2h&poolSessionTimeout=2h&poolIncrement=30&timezone=Local&connect_timeout=15"

	// 自动拼接成 select * from marvin.marvin
	oraConf.SchemaName = "marvin"
	oraConf.Table = "marvin"

	ctx := context.TODO()

	if err := NewReverser(ctx, oraConf, mysqlConf).Reverse(); err != nil {
		fmt.Printf("\nreverse failed: %s\n", err.Error())
		os.Exit(1)
	}

}
