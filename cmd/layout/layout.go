package layout

import (
	"github.com/spf13/cobra"
	"github.com/wangyufengGoGoGo/generate-project-layout/pkg/folder"
	"github.com/wangyufengGoGoGo/generate-project-layout/pkg/global"
)

/**
 * @Author wyf
 * @Date 2021/10/19 14:53
 **/

var Cmd = &cobra.Command{
	Use: "new", Short: "generate project layout",
	Run: func(cmd *cobra.Command, args []string) {
	folder.GenerateDIr(global.Layout)
}}
