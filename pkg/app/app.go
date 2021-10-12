package app

import (
	"flag"
	"fmt"
	"github.com/wangyufengGoGoGo/generate-project-layout/pkg/folder"
	"github.com/wangyufengGoGoGo/generate-project-layout/pkg/global"
)

func NewApp() error {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return fmt.Errorf("参数错误")
	}

	switch args[0] {
	case "new":
		if err := folder.GenerateDIr(global.Layout); err != nil {
			return err
		}
	default:
		return fmt.Errorf("参数不存在")
	}

	return nil
}
