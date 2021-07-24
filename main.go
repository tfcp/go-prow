package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                      // ui theme
	"go-prow/bootstrap"
)

func main() {
	bootstrap.StartServer()
}
