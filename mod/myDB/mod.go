package myDB

import (
	"errors"
	"fmt"
	"github.com/juanjiTech/jframe/conf"
	"github.com/juanjiTech/jframe/core/kernel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ kernel.Module = (*Mod)(nil)

type Mod struct {
	kernel.UnimplementedModule // 请为所有Module引入UnimplementedModule
}

func (m *Mod) Name() string {
	return "myDB"
}

func (m *Mod) PreInit(hub *kernel.Hub) error {
	c := conf.Get().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s"+
		"?charset=utf8mb4&parseTime=True&loc=Local&tls=%v",
		c.USER, c.PASSWORD, c.Addr, c.PORT, c.DATABASE, c.UseTLS)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	hub.Log.Info("mysql init success")
	hub.Map(&db)
	return nil
}

func (m *Mod) Init(hub *kernel.Hub) error {
	// check if inject success
	var db *gorm.DB
	if hub.Load(&db) != nil {
		return errors.New("can't load gorm from kernel")
	}
	result := db.Exec("SHOW TABLES")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
