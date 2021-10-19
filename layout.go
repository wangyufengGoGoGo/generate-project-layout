package main

import (
	"github.com/spf13/cobra"
	"github.com/wangyufengGoGoGo/generate-project-layout/cmd/layout"
)

func main() {
	root := &cobra.Command{
		Use:   "layout",
		Short: "generate project layout",
	}
	root.AddCommand(
		layout.Cmd,
		)
	root.Execute()
}
