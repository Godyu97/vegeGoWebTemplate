package dao

import (
	"context"
	"github.com/Godyu97/vege9/vegedsn"
	"github.com/Godyu97/vegeGoWebTemplate/internal/config"
	"github.com/Godyu97/vegeGoWebTemplate/internal/dal/query"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// all db obj
var mdb *gorm.DB
var rdb *redis.Client

// InitDb in project main init
func InitDb(m config.DSN, r config.RedisConf) {
	var err error
	//mysql
	mdb, err = gorm.Open(
		mysql.Open(
			vegedsn.NewDefaultMysqlDsn(
				vegedsn.DefaultParams,
				vegedsn.WithAuth(m.Username, m.Password),
				vegedsn.WithAddress(m.Address),
				vegedsn.WithDatabase(m.Database),
			).String()),
		&gorm.Config{
			Logger: logger.Default,
		})
	if err != nil {
		panic(err)
	}
	SetDBMaxConns(mdb)
	//gen query DB Set
	queryInit()
	//redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.Database,
	})
}

func SetDBMaxConns(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(25)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
}

func GetMDBMySql() *gorm.DB {
	if mdb != nil {
		return mdb
	}
	panic("Init Mysql First")
}

func GetRedis() *redis.Client {
	if rdb != nil {
		return rdb
	}
	panic("Init Redis First")
}

func CloseDbs() {
	GetRedis().Close()
	sqlDB, _ := GetMDBMySql().DB()
	sqlDB.Close()
}

var qmdb *query.Query

func queryInit() {
	qmdb = query.Use(mdb)
}

func GetMQuery() *query.Query {
	if qmdb != nil {
		return qmdb
	}
	panic("Init Query First")
}

var bgCtx = context.Background()
