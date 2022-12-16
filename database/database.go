package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

type (
	connection struct{}

	SqlLogger struct {
		logger.Interface
	}
)

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, _ := fc()
	fmt.Printf("\n => SQL Statement :: %v \n\n", sqlStatement)
}

func NewDatabase() *connection {
	return &connection{}
}

func (conn *connection) MySQLInit() *sql.DB {
	db, err := conn.mySqlConnection()
	if err != nil {
		panic(err)
	}

	return db
}

func (conn *connection) GormInit() *gorm.DB {
	db, err := conn.gormMySqlConnection()
	if err != nil {
		panic(err)
	}

	return db
}

func (conn *connection) RedisInit() *redis.Client {
	return conn.redisConnection()
}

func (conn *connection) mySqlConnection() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", viper.GetString("sql-db.username"), viper.GetString("sql-db.password"), viper.GetString("sql-db.hostname"), viper.GetString("sql-db.dbName"))
	driverName := viper.GetString("sql-db.driver-name")

	return sql.Open(driverName, dsn)

}

func (conn *connection) gormMySqlConnection() (*gorm.DB, error) {

	// * ------------ ใช้ *sql.DB ในการเชื่อมต่อกับ GORM ---------------
	sqlDB, err := conn.mySqlConnection()
	if err != nil {
		panic(err)
	}

	return gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: &SqlLogger{},
	})
}

func (conn *connection) redisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis-cache.address"),
		Password: viper.GetString("redis-cache.password"),
		DB:       viper.GetInt("redis-cache.db-num"),
	})
}
