package controller

import (
	"tuhla/common/db"
)

type Controller struct {
	db *db.DBConnection
}

func New(db *db.DBConnection) *Controller {
	return &Controller{db}
}
