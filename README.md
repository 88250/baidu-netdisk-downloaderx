## BND

BND（Baidu Netdisk Downloader）是一个用 golang 编写的百度网盘不限速下载器。

项目地址 https://github.com/b3log/baidu-netdisk-downloaderx 好用记得给颗星！

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

![BND 教学](https://img.hacpai.com/file/2018/05/c87225b75d12411ca5ec4a57274371eb_.gif)

### 常见问题

#### 如何获取 BDUSS？

请参考[这里](https://www.baidu.com/s?wd=如何获取BDUSS)。

#### 如何下载非根目录的文件（文件夹）？

请在百度网盘里把要下载的文件移动到根目录，然后在 BND 中刷新文件列表。

#### 速度忽高忽低？

BND 显示的速度是瞬时速度，仅作为参考，不必在太在意。

#### 360 报毒？

这个我也不知道为啥，但我敢拍着胸脯保证无毒无害，业界良心！

#### Mac/Linux 下如何粘贴 BDUSS？

Mac/Linux 下在输入 BDUSS 时请使用鼠标右键的下拉菜单，不支持快捷键粘贴。

### 问题反馈

* [Q 群 739075568](https://shang.qq.com/wpa/qunwpa?idkey=e1b4287d075e86792f42f413f75943c91da37d074649d28c51aa6d48361631ba)
* [论坛](https://hacpai.com/article/1524460877352)

### 更新日志

请看这里[这里](https://github.com/b3log/baidu-netdisk-downloaderx/blob/master/CHANGE_LOGS.md)。

### 鸣谢

* [BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)：百度网盘客户端 - Go 语言编写
* [andlabs/ui](https://github.com/andlabs/ui)：跨平台的 Go GUI 库
