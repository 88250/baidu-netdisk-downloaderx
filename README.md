## 百度网盘不限速下载器

又一个百度网盘不限速下载器，项目地址 https://github.com/b3log/baidu-netdisk-downloaderx 好用记得给颗星！

### 特性

* 简单友好的图形用户界面
* 支持 Windows、Mac、Linux
* 通过 Cookie [BDUSS] 登录，无需担心密码泄漏
* 多线程下载
* 支持断点续传

![主界面](https://img.hacpai.com/file/2018/04/5aebc46de06c4d29aec91d65751aff5a_.png)

### 用法

1. [下载发布包](https://share.weiyun.com/57zViCm)
2. 在浏览器登录百度网盘，获取 BDUSS Cookie 值（**注意拷贝完整**）
   ![BDUSSpng](https://img.hacpai.com/file/2018/04/d1a78d5163f644d7931925ef5edbf9dd_BDUSS.png)
3. 登录成功后显示主界面，选择需要下载的文件后点击下载即可

### 说明

* Mac 和 Linux 下在输入 BDUSS 时请使用鼠标右键的下拉菜单，不支持快捷键粘贴
* 程序只会获取根目录下最新修改的 10 个文件，请将需要下载的文件放到根目录下
* 不支持 Windows XP，以及 32 位的操作系统
* 下载的数据默认保存在用户的 Downloads 文件夹下

### 反馈

* [Q 群 739075568](https://shang.qq.com/wpa/qunwpa?idkey=e1b4287d075e86792f42f413f75943c91da37d074649d28c51aa6d48361631ba)
* [论坛](https://hacpai.com/article/1524460877352)

### 更新日志

#### v2.1.0 / 2018-04-28

* Windows 版下载路径可供用户选择

#### v2.0.4 / 2018-04-24

* 调整界面细节
* 下载路径默认用户目录下的 Downloads 文件夹
* 发布[吾爱破解论坛特别版](https://www.52pojie.cn/thread-730453-1-1.html)
* 发布[黑客派论坛特别版](https://hacpai.com/article/1524460877352)

#### v2.0.3 / 2018-04-21

第一次公开发布。

### 鸣谢

* [BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)：百度网盘客户端 - Go 语言编写
* [andlabs/ui](https://github.com/andlabs/ui)：跨平台的 Go GUI 库
