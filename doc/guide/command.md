# 命令

## 指南

colon 允许自定义命令，挂载到 colon 上。并且提供了`./colon command` 系列命令。

```
[~/go-colon/colon/demo]$ ./colon command
all about commond

Usage:
  colon command [flags]
  colon command [command]

Available Commands:
  list        show all command list
  new         create a command

Flags:
  -h, --help   help for command

Use "colon command [command] --help" for more information about a command.
```

- list  // 列出当前所有已经挂载的命令列表
- new   // 创建一个新的自定义命令

## 创建

创建一个新命令，可以使用 `./colon command new`

这是一个交互式的命令行工具。

```
[~/go-colon/colon/demo]$ ./colon command new
create a new command...
? please input command name: test
? please input file name(default: command name):
create new command success，file path: ~/go-colon/colon/demo/app/console/command/test.go
please remember add command to console/kernel.go
```

创建完成之后，会在应用的 app/console/command/ 目录下创建一个新的文件。

## 自定义

colon 中的命令使用的是 cobra 库。 https://github.com/spf13/cobra

```
package command

import (
        "fmt"

        "github.com/go-colon/colon/framework/cobra"
        "github.com/go-colon/colon/framework/command/util"
)

var TestCommand = &cobra.Command{
        Use:   "test",
        Short: "test",
        RunE: func(c *cobra.Command, args []string) error {
                container := util.GetContainer(c.Root())
                fmt.Println(container)
                return nil
        },
}

```

基本上，我们要求实现
- Use // 命令行的关键字
- Short // 命令行的简短说明
- RunE // 命令行实际运行的程序

更多配置和参数可以参考 [cobra 的 github 页面](https://github.com/spf13/cobra)

## 挂载

编写完自定义命令后，请记得挂载到 console/kernel.go 中。

``` golang
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "colon",
		Short: "main",
		Long:  "colon commands",
	}

	ctx := commandUtil.RegiestContainer(container, rootCmd)

	colonCommand.AddKernelCommands(rootCmd)

    rootCmd.AddCommand(command.DemoCommand)

	return rootCmd.ExecuteContext(ctx)
}

```
