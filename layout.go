package main

import (
	"github.com/wangyufengGoGoGo/generate-project-layout/pkg/app"
)

func main() {
	if err := app.NewApp(); err != nil {
		panic(err)
	}
}
