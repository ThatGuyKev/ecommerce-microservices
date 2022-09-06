package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// shardingMiddleware := sharding.Register(sharding.Config{
	// 	ShardingKey:         "target_demographic_context_id",
	// 	NumberOfShards:      50,
	// 	PrimaryKeyGenerator: sharding.PKSnowflake,
	// }, "products")

	// db.Use(shardingMiddleware)

	// db.Exec("CREATE TABLE 'product' (id NOT NULL, context_id VARCHAR(36) NOT NULL")

	return Handler{db}
}
