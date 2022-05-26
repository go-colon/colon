# swagger

## 命令

colon 使用 [swaggo](https://github.com/swaggo/swag) 集成了 swagger 生成和服务项目。
并且封装了 `./colon swagger` 命令。

```
[~/go-colon/colon/demo]$ ./colon swagger
swagger对应命令

Usage:
  Colon swagger [flags]
  Colon swagger [command]

Available Commands:
  gen         生成对应的swagger文件, contain swagger.yaml, doc.go

Flags:
  -h, --help   help for swagger

Use "Colon swagger [command] --help" for more information about a command.

```

- gen  // 生成swagger文件

待实现：
- serve // 提供swagger服务

## 注释

colon 使用 [swaggo](https://github.com/swaggo/swag) 来实现注释生成 swagger 功能。

全局注释在文件  `app/http/swagger.go` 中

接口注释请写在各自模块的 api.go 中

```
// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /user/all [get]
func (api *UserApi) All(c *gin.Context) {
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}
```

swagger 注释的格式和关键词可以参考：[swaggo](https://github.com/swaggo/swag)

## 生成

使用命令 `./colon swagger gen`

```
[~/go-colon/colon/demo]$ ./colon swagger gen
2022/05/26 17:43:54 Generate swagger docs....
2022/05/26 17:43:54 Generate general API Info, search dir:/Users/cb/go/src/yuewen/colon/app/http
2022/05/26 17:43:54 Generating demo.UserDTO
2022/05/26 17:43:54 Generating user.loginParam
2022/05/26 17:43:54 Generating user.registerParam
2022/05/26 17:43:54 create docs.go at /Users/cb/go/src/yuewen/colon/app/http/swagger/docs.go
2022/05/26 17:43:54 create swagger.json at /Users/cb/go/src/yuewen/colon/app/http/swagger/swagger.json
2022/05/26 17:43:54 create swagger.yaml at /Users/cb/go/src/yuewen/colon/app/http/swagger/swagger.yaml
```

在目录 `app/http/swagger/` 下自动生成swagger相关文件。

## 服务

可以使用命令 `./colon swagger serve` 启动当前应用的 swagger ui 服务。

```
[~/go-colon/colon/demo]$ ./colon swagger serve
swagger success: http://0.0.0.0:8069/swagger/index.html
if you want to replace, remember use command: swagger gen
```

>
> 如果你的 swagger 服务已经启动，更新 swagger 只需要重新运行 `./colon swagger gen` 就能更新。
> 因为 swagger 服务读取的是生成的 swagger.json 这个文件。


服务端口，我们也可以通过配置文件 `config/[env]/swagger.yaml` 中的配置
```
url: http://127.0.0.1:8069
```

来配置swagger serve 启动的服务。
