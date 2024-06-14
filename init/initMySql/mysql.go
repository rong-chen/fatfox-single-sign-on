package initMySql

import (
	"database/sql"
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/init/initViper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMySQL(config *initViper.Config) {
	// 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.MYSQL.User, config.MYSQL.Password, config.MYSQL.Host, config.MYSQL.Port)
	db, err := sql.Open("mysql", dsn)
	dbName := "registrationCenter"
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 检查数据库是否存在
	exists, err := databaseExists(db, dbName)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		err = createDatabase(db, dbName)
		if err != nil {
			log.Fatal(err)
		}
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	global.MySqlDb, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MYSQL.User, config.MYSQL.Password, config.MYSQL.Host, config.MYSQL.Port, config.MYSQL.Db)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("连接数据库失败，请检查参数:", err)
	}
}

// 检查数据库是否存在
func databaseExists(db *sql.DB, dbName string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = '%s'", dbName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 创建数据库
func createDatabase(db *sql.DB, dbName string) error {
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err := db.Exec(query)
	return err
}
