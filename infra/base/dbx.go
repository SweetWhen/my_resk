package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"myresk/infra"
)

var database * dbx.Database

func DbxDatabase() *dbx.Database  {

	return database
}

//dbx数据库starter,并且设置为全局
type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter)Setup(ctx infra.StarterContext)  {
	conf := ctx.Props()

	settings := dbx.Settings{
		//DriverName:"mysql",
		//User:"po",
		//Password:"casa123",
		//Database:"po0",
		//Host: "127.0.0.1:3306",
	}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Infof("settings:%+v\n", settings)
	logrus.Info("mysql conn url: ",settings.ShortDataSourceName())
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}