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
