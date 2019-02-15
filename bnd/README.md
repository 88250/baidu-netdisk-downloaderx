## BND

[又一个百度网盘不限速下载器 BND](https://hacpai.com/article/1524460877352)

### 编译

1. 安装 golang 环境
2. 参考 https://github.com/andlabs/libui 编译 UI 库
3. 不支持交叉编译，只能在目标平台上编译
4. Windows 执行 build.bat。其他平台可参考该脚本进行构建

### 其他

* aria2 原有设计是在启动后检查版本并远程拉取的，现已改为本地打包
* 有保留了版本检查机制，可搜索 rhythm.b3log.org 进行相关修改
* 和服务端交互时用于加密请求响应数据的密钥已在源码中公开