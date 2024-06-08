package model

import (
	"github.com/lgstvo/crud-golang/src/configuration/logger"
	"github.com/lgstvo/crud-golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}
