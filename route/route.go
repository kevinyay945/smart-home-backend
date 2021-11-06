package route

import (
	"smart-home-backend/model"
	v1 "smart-home-backend/route/v1"
)

func NewVersion1() v1.IRoute {
	request := model.NewRequest()
	command := model.NewCommand()
	return v1.New(request, command)
}

