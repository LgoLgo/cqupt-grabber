# 重庆邮电大学抢课工具

## 目录
- [重庆邮电大学抢课框架](#重庆邮电大学抢课框架)
    - [目录](#目录)
    - [特别声明](#特别声明)
    - [安装](#安装)
    - [快速开始](#快速开始)
    - [COOKIE以及LOADS的获取教程](#COOKIE以及LOADS的获取教程)
         - [COOKIE](#COOKIE)
         - [LOAD](#LOAD)
    - [其余高级操作](#其余高级操作)

## 特别声明

- 本仓库发布的脚本及其中涉及的任何功能，仅用于测试和学习研究，禁止用于商业用途，禁止用于违法用途，不能保证其准确性，完整性和有效性，请根据情况自行判断。

- 本项目内所有资源文件，禁止任何公众号、自媒体进行任何形式的转载、发布。

- 本人对任何脚本问题概不负责，包括但不限于由任何脚本错误导致的任何损失或损害。

- 请勿将本仓库的任何内容用于商业或非法目的，否则后果自负。

- 如果任何单位或个人认为该项目的脚本可能涉嫌侵犯其权利，则应及时通知并提供身份证明，所有权证明，我将在收到认证文件后删除相关脚本。

- 任何以任何方式查看此项目的人或直接或间接使用该项目的任何脚本的使用者都应仔细阅读此声明。本人保留随时更改或补充此免责声明的权利。一旦使用并复制了任何相关脚本或Go项目的代码，则视为您已接受此免责声明。

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
使用包中的LoopRob，0.25s进行一次抢课，直到有一门课被抢到
```go
package main

import (
	"github.com/L2ncE/CQUPT-ClassGrabbing/ClassGrabbing"
)

func main() {
	cookie := "这里是一个cookie"
	
	//支持同时抢多门课程
	loads := []string{
		"这里是第一节课",
		"这里是第二节课"}

	ClassGrabbing.LoopRob(cookie, loads)
}
```
其中cookie以及loads需要自己获取

## COOKIE以及LOADS的获取教程

### COOKIE
首先进入到[选课系统](http://xk1.cqupt.edu.cn/)中登录进入到选课详细界面

在键盘上按下**F12**进入开发者工具并点击网络选项卡

![image-20220513233644944](https://s2.loli.net/2022/05/13/sM4mAclvHnuyO2F.png)

在键盘上按**F5**进行页面刷新后，网络视图同样会进行刷新，我们点击yxk.php后选择标头选项卡

![image-20220513233828066](https://s2.loli.net/2022/05/13/gc8bBAzf1qwevUr.png)

然后一直往下翻，就会找到**COOKIE**，将其复制下来即可

![image-20220513233921962](https://s2.loli.net/2022/05/13/czhZMt2aL5U1wup.png)

### LOAD

#### 方法一

重复和上面COOKIE的相同步骤到选课界面，在你想选择的课程旁边点击“+”号，查看网络选项卡会新增一个POST请求，将其负载选项卡中的源代码复制下来即可

![image-20220513234236467](https://s2.loli.net/2022/05/13/H6xGfKPQ9d5aVeN.png)



#### 方法二

我们的抢课工具中就已经封装了获取load的函数

> 分别可以获取所有人文、自然选课的负载（会有课程相关的信息提示），除此以外还可以通过模糊搜索进行快速准确的查找

```go
package main

import (
	"github.com/L2ncE/CQUPT-ClassGrabbing/Query"
)

func main() {
	cookie := "这里是一个cookie"
	param := "Rw or Zr" //其中Rw为人文选修，Zr为自然选修
	content := "你想模糊搜索的内容" //例如输入“工程”会将所有带有工程两个字的课程信息以及负载输出

	Query.AllRenWen(cookie)
	Query.AllZiRan(cookie)
	Query.Search(param, cookie, content)
}

```

## 其余高级操作

```go
//高并发抢课 会有被BAN风险
func LoopRobWithHighConcurrency(cookie string, loads []string) {
...
}
...
```

```go
//只进行一次访问并传回响应
func SingleRobWithInfo(cookie string, load string) {
...
}
...
```

```go
//自定义一次访问的速度
//duration中为你想自定义的秒数，建议不小于0.2
func LoopRobWithCustomTime(cookie string, loads []string, duration float64) {
...
}
...
```

