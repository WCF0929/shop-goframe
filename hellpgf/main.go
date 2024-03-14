package main

import (
	_ "hellpgf/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" //初始化

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "hellpgf/internal/packed"
	//
	"hellpgf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
