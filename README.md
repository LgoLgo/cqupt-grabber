# 重庆邮电大学抢课框架

## 目录
- [重庆邮电大学抢课框架](#重庆邮电大学抢课框架)
    - [目录](#目录)
    - [特别声明](#特别声明)
    - [安装](#安装)
    - [快速开始](#快速开始)
    - [Build with jsoniter/go-json](#build-with-json-replacement)
    - [Build without `MsgPack` rendering feature](#build-without-msgpack-rendering-feature)
    - [API Examples](#api-examples)
        - [Using GET, POST, PUT, PATCH, DELETE and OPTIONS](#using-get-post-put-patch-delete-and-options)
        - [Parameters in path](#parameters-in-path)
        - [Querystring parameters](#querystring-parameters)
        - [Multipart/Urlencoded Form](#multiparturlencoded-form)
        - [Another example: query + post form](#another-example-query--post-form)
        - [Map as querystring or postform parameters](#map-as-querystring-or-postform-parameters)
        - [Upload files](#upload-files)
            - [Single file](#single-file)
            - [Multiple files](#multiple-files)

## 特别声明

- 本仓库发布的脚本及其中涉及的任何功能，仅用于测试和学习研究，禁止用于商业用途，不能保证其合法性，准确性，完整性和有效性，请根据情况自行判断。

- 本项目内所有资源文件，禁止任何公众号、自媒体进行任何形式的转载、发布。

- 本人对任何脚本问题概不负责，包括但不限于由任何脚本错误导致的任何损失或损害。

- 请勿将本仓库的任何内容用于商业或非法目的，否则后果自负。

- 如果任何单位或个人认为该项目的脚本可能涉嫌侵犯其权利，则应及时通知并提供身份证明，所有权证明，我将在收到认证文件后删除相关脚本。

- 任何以任何方式查看此项目的人或直接或间接使用该项目的任何脚本的使用者都应仔细阅读此声明。本人保留随时更改或补充此免责声明的权利。一旦使用并复制了任何相关脚本或Script项目的规则，则视为您已接受此免责声明。

**您必须在下载后的24小时内从计算机或手机中完全删除以上内容**

> ***您使用或者复制了本仓库的脚本，则视为 `已接受` 此声明，请仔细阅读***

## 安装

我将默认你已经拥有Go语言开发环境，如果你还没有可以去到[这个链接](https://blog.csdn.net/weixin_44621343/article/details/117792504)

1. 你需要Go的开发环境[Go](https://golang.org/)，然后您可以使用以下 Go 命令安装 Gin
```sh
$ go get -u github.com/L2ncE/CQUPT-ClassGrabbing
```

2. 在你的代码中import

```go
import "github.com/L2ncE/CQUPT-ClassGrabbing"
```


## 快速开始
使用包中的LoopRob
```go
package main

import (
	"github.com/L2ncE/CQUPT-ClassGrabbing/classRobbing"
)

func main() {
	cookie := "*********"
	loads := []string{"*****",
		"*****"}
	
	classRobbing.LoopRob(cookie, loads)
}
```
其中cookie以及loads需要自己获取

