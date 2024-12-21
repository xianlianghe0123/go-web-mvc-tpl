package server

import (
	"context"
	"os"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/repository/query"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

type Migrate struct {
	db  *query.Query
	log *log.Logger
}

func NewMigrate(db *query.Query, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}

func (m *Migrate) Start(ctx context.Context) error {
	m.log.Info("AutoMigrate success")
	os.Exit(0)
	return nil
}

func (m *Migrate) Stop(ctx context.Context) error {
	m.log.Info("AutoMigrate stop")
	return nil
}
