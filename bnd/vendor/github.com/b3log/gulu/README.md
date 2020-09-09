<p align = "center">
<img alt="Wide" src="https://user-images.githubusercontent.com/873584/58315007-4100f080-7e43-11e9-9b10-b64a6a4a5d2d.png">
<br><br>
Go 语言常用工具库，这个轱辘还算圆！
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/b3log/gulu"><img src="https://img.shields.io/travis/b3log/gulu.svg?style=flat-square"></a>
<a title="GoDoc" target="_blank" href="https://godoc.org/github.com/b3log/gulu"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/b3log/gulu"><img src="https://goreportcard.com/badge/github.com/b3log/gulu?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/b3log/gulu/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/b3log/gulu.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/b3log/gulu"><img src="https://img.shields.io/github/languages/code-size/b3log/gulu.svg?style=flat-square"></a>
<br>
<a title="Apache License" target="_blank" href="https://github.com/b3log/gulu/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-apache2-orange.svg?style=flat-square"></a>
<a title="GitHub Commits" target="_blank" href="https://github.com/b3log/gulu/commits/master"><img src="https://img.shields.io/github/commit-activity/m/b3log/gulu.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/b3log/gulu/commits/master"><img src="https://img.shields.io/github/last-commit/b3log/gulu.svg?style=flat-square&color=FF9900"></a>
<a title="GitHub Pull Requests" target="_blank" href="https://github.com/b3log/gulu/pulls"><img src="https://img.shields.io/github/issues-pr-closed/b3log/gulu.svg?style=flat-square&color=FF9966"></a>
<a title="Hits" target="_blank" href="https://github.com/b3log/hits"><img src="https://hits.b3log.org/b3log/gulu.svg"></a>
<br><br>
<a title="GitHub Watchers" target="_blank" href="https://github.com/b3log/gulu/watchers"><img src="https://img.shields.io/github/watchers/b3log/gulu.svg?label=Watchers&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Stars" target="_blank" href="https://github.com/b3log/gulu/stargazers"><img src="https://img.shields.io/github/stars/b3log/gulu.svg?label=Stars&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Forks" target="_blank" href="https://github.com/b3log/gulu/network/members"><img src="https://img.shields.io/github/forks/b3log/gulu.svg?label=Forks&style=social"></a>&nbsp;&nbsp;
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>
</p>

## ✨ 功能

<details>
<summary>文件操作 <code>gulu.File</code></summary>
<br>

* 获取文件大小
* 判断路径是否存在
* 判断文件是否是图片
* 按内容判断文件是否是可执行二进制
* 判断文件是否是目录
* 复制文件
* 复制目录
</details>

<details>
<summary>Go 语言 <code>gulu.Go</code></summary>
<br>

* 获取 Go API 源码目录路径
* 判断指定路径是否在 Go API 源码目录下
* 获取格式化工具名 ["gofmt", "goimports"]
* 获取 $GOBIN 下指定可执行程序名的绝对路径
</details>


<details>
<summary>日志记录 <code>gulu.Log</code></summary>
<br>

* 提供可指定日志级别的日志记录器
</details>

<details>
<summary>网络相关 <code>gulu.Net</code></summary>
<br>

* 获取本机第一张网卡的地址
</details>

<details>
<summary>操作系统 <code>gulu.OS</code></summary>
<br>

* 判断是否是 Windows
* 获取当前进程的工作目录
* 获取用户 Home 目录路径
</details>

<details>
<summary>panic 处理 <code>gulu.Panic</code></summary>
<br>

* 包装 recover() 提供更好的报错日志格式
</details>

<details>
<summary>随机数 <code>gulu.Rand</code></summary>
<br>

* 随机字符串
* 随机整数
</details>

<details>
<summary>返回值相关 <code>gulu.Ret</code></summary>
<br>

* 提供普适返回值结构
</details>

<details>
<summary>Rune 相关 <code>gulu.Rune</code></summary>
<br>

* 判断 rune 是否为数字或字母
* 判断 rune 是否为字母
</details>

<details>
<summary>字符串相关 <code>gulu.Str</code></summary>
<br>

* 字符串是否包含在字符串数组中
* 求最长公共子串
</details>

<details>
<summary>Zip 压缩解压<code>gulu.Zip</code></summary>
<br>

* Zip 压缩和解压
</details>

## 🗃 案例

* [Pipe](https://github.com/b3log/pipe)：一款小而美的博客平台，专为程序员设计
* [Wide](https://github.com/b3log/wide)：一款基于 Web 的 Go 语言 IDE，随时随地玩 golang
* [协慌网](https://routinepanic.com)：专注编程问答汉化

如果你也在使用 Gulu，欢迎通过 PR 将你的项目添加到这里。

## 💝 贡献

Gulu 肯定有一些不足之处：

* 实现存在缺陷
* 代码不够优美
* 文档不够清晰
* 功能不够完善
* ……

希望大家能和我们一起来完善该项目，无论是提交需求建议还是代码改进，我们都非常欢迎！

## 🏘️ 社区

* [讨论区](https://ld246.com/tag/gulu)
* [报告问题](https://github.com/b3log/gulu/issues/new/choose)

## 📄 授权

Gulu 使用 [Apache License, Version 2](https://www.apache.org/licenses/LICENSE-2.0) 开源协议。

## 🙏 鸣谢

* [The Go Programming Language](https://golang.org)
