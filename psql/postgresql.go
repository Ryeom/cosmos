package psql

import (
	"errors"
	"github.com/Ryeom/cosmos/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var Postgresql DataBase

type DataBase struct {
	database *gorm.DB
	dsn      string
}

func InitializePostgresql() {
	err := newDatabaseSession()
	if err != nil {
		panic(err)
	}
}

func newDatabaseSession() error {
	dsn := viper.GetString("postgresql.key")
	if os.Args[1] == "local" || os.Args[1] == "dev" {
		log.Logger.Info("DB Information : ", dsn)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true, // 단수 테이블 이름을 사용합니다. `User`의 테이블은 이 옵션을 활성화하면 `user`가 됩니다.
		},
		DisableAutomaticPing: false, // 자동 핑
	})

	if err != nil {
		return errors.New("Fail Database Connection." + err.Error())
	}
	Postgresql = DataBase{
		database: db,
		dsn:      dsn,
	}

	return nil
}
func (d *DataBase) GetDB() *gorm.DB {
	if d.database == nil {
		err := newDatabaseSession()
		if err != nil {
			// TODO noti
		}
	}
	//fmt.Println(d.database.Statement.Schema.String())
	return d.database
}
