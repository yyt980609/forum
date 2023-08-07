package service

import (
	"forum/dao"
	"forum/utils/snowflake"

	"go.uber.org/zap"
)

func Sign() {
	dao.QueryUser()
	id, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("Generate id failed.")
		return
	}
	println(id)
	dao.InsertUser()
}
