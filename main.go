package main

import (
	config2 "github.com/springeye/droidjs-cli/config"
	"github.com/teris-io/cli"
	"os"
)

func main() {
	list := cli.NewCommand("list", "查看云端文件列表").
		WithAction(func(args []string, options map[string]string) int {
			// do something
			return 0
		})

	pull := cli.NewCommand("pull", "从云端拉取一个文件到本地").
		WithArg(
			cli.NewArg("filename", "文件名").
				AsOptional().
				WithType(cli.TypeString),
		).
		WithOption(
			cli.NewOption("interactive", "交互模式").
				WithChar('i').
				WithType(cli.TypeBool),
		).
		WithAction(func(args []string, options map[string]string) int {
			// do something
			return 0
		})

	config := cli.NewCommand("config", "配置客户端").
		WithShortcut("cfg").
		WithOption(cli.NewOption("edit", "修改现有的配置").WithChar('e').WithType(cli.TypeBool)).
		WithAction(config2.SetupConfig)
	app := cli.New("git tool").
		WithOption(cli.NewOption("verbose", "Verbose execution").WithChar('v').WithType(cli.TypeBool)).
		WithCommand(list).
		WithCommand(pull).
		WithCommand(config)

	// no action attached, just print usage when executed

	os.Exit(app.Run(os.Args, os.Stdout))
}
