package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

// NewDB 返回 *gorm.DB
func NewDB(c Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName:           c.Driver,
		DSN:                  buildDNS(c),
		PreferSimpleProtocol: c.PreferSimpleProtocol,
		Conn:                 c.Conn,
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   c.Logger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if c.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	}

	if c.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	}

	if c.ConnMaxLifeTime > 0 {
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifeTime)
	}

	return db, nil
}

// buildDNS 构建连接数据库的 dns
func buildDNS(c Config) string {
	options := strings.Join(c.Options, " ")
	dsn := "host=" + c.Host + " port=" + c.Port + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.Database + " " + options
	return dsn
}
