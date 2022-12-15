package main

import (
	"database/sql"

	"demo-mysql-gorm/config"
	"demo-mysql-gorm/database"
	"demo-mysql-gorm/server"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var mysql_db *sql.DB
var gorm_db *gorm.DB
var redis_cache *redis.Client

func init() {

	config.ConfigInit()
	mysql_db = database.NewDatabase().MySQLInit()
	redis_cache = database.NewDatabase().RedisInit()
	gorm_db = database.NewDatabase().GormInit()

}

func main() {

	defer mysql_db.Close()
	defer redis_cache.Close()

	app := fiber.New()

	server := server.NewServer(mysql_db, gorm_db, redis_cache)
	server.Router(app)

	app.Listen(viper.GetString("demo.port"))

}
