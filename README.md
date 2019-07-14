## 简介

<a title="Hits" target="_blank" href="https://github.com/b3log/hits"><img src="https://hits.b3log.org/b3log/baidu-netdisk-downloaderx.svg"></a>
<a title="GitHub Watchers" target="_blank" href="https://github.com/b3log/baidu-netdisk-downloaderx/watchers"><img src="https://img.shields.io/github/watchers/b3log/baidu-netdisk-downloaderx.svg?label=Watchers&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Stars" target="_blank" href="https://github.com/b3log/baidu-netdisk-downloaderx/stargazers"><img src="https://img.shields.io/github/stars/b3log/baidu-netdisk-downloaderx.svg?label=Stars&style=social"></a>&nbsp;&nbsp;
<a title="GitHub Forks" target="_blank" href="https://github.com/b3log/baidu-netdisk-downloaderx/network/members"><img src="https://img.shields.io/github/forks/b3log/baidu-netdisk-downloaderx.svg?label=Forks&style=social"></a>&nbsp;&nbsp;
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>

[BND](https://github.com/b3log/baidu-netdisk-downloaderx)（Baidu Netdisk Downloader）是一款图形界面的百度网盘不限速下载器，支持 Windows、Linux、Mac。

BND 分为两个系列，BND1 和 BND2，下面分别进行介绍。

## BND1

[又一个百度网盘不限速下载器 BND](https://hacpai.com/article/1524460877352)

![bnd1](https://user-images.githubusercontent.com/873584/52854939-0825f100-315b-11e9-9aca-d03841b6c44e.png)

* 小巧省资源
* 支持 Linux、支持 32 位的 Windows

### 代码

本项目是基于 [BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go) 开发：

* 在其基础上增加了 UI 界面，主要修改点是 pcscommand 包
* Windows 版引入了 Aria2，下载超过 512M 文件时会切换到 Aria2

### 编译

1. 安装 golang 环境
2. 项目目录 $GOPATH/src/github.com/b3log/bnd
3. 参考 https://github.com/andlabs/libui 编译 UI 库
4. 不支持交叉编译，只能在目标平台上编译
5. Windows 执行 build.bat，Linux/Mac 执行 build.sh

### 其他

* aria2 原有设计是在启动后检查版本并远程拉取的，现已改为本地打包
* 保留了版本检查机制，可搜索 rhythm.b3log.org 进行相关修改
* 和服务端交互时用于加密请求响应数据的密钥已在源码中公开

### 鸣谢

* [aria2](https://github.com/aria2/aria2)：超高速的下载引擎
* [BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)：百度网盘客户端 - Go 语言编写
* [andlabs/ui](https://github.com/andlabs/ui)：跨平台的 Go GUI 库

## BND2

[百度不限速下载器 BND2 技术架构简介](https://hacpai.com/article/1535277215816)

![bnd2](https://user-images.githubusercontent.com/970828/52894072-df473f80-31de-11e9-8301-33d2fa9858b4.png)

* 界面美观，操作便捷
* 支持多任务并发下载

### 编译

1. 安装 golang、node 环境
2. 项目目录 $GOPATH/src/github.com/b3log/bnd2
3. Windows 执行 build.bat
4. Mac 在 electron 目录下执行 `npm install && npm run dist`
5. `electron/dist` 目录下运行可执行文件

### 其他

* 内核可执行文件以及 aria2 原有设计是在启动后检查版本并远程拉取的，现已改为本地打包
* 保留了版本检查机制，可搜索 rhythm.b3log.org 进行相关修改
* 和服务端交互时用于加密请求响应数据的密钥已在源码中公开

### 鸣谢

* [aria2](https://github.com/aria2/aria2)：超高速的下载引擎
* [React](https://github.com/facebook/react)：使用 JS 构建用户界面库
* [Electron](https://github.com/electron/electron)：使用 JS、HTML、CSS 的跨平台桌面应用库

## 贡献者

BND 项目的主要贡献者来自于 B3log 开源社区，大家对 BND 的开发、测试、反馈、推广等贡献了自己的一份力量。

[B3log 开源组织欢迎大家加入！](https://hacpai.com/article/1463025124998)

---

## 开源项目推荐

* 如果你需要搭建一个个人博客系统，可以考虑使用 [Solo](https://github.com/b3log/solo)
* 如果你需要搭建一个多用户博客平台，可以考虑使用 [Pipe](https://github.com/b3log/pipe)
* 如果你需要搭建一个社区平台，可以考虑使用 [Sym](https://github.com/b3log/symphony)
* 欢迎加入我们的小众开源社区，详情请看[这里](https://hacpai.com/article/1463025124998)
