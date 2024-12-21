package service

import (
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/repository"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
