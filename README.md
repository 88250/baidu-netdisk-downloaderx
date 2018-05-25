## BND

BND（Baidu Netdisk Downloader）是一款图形界面的百度网盘不限速下载器，支持 Windows、Linux、Mac。

BND 由 [B3log 开源社区](https://github.com/b3log)负责维护，项目地址 https://github.com/b3log/baidu-netdisk-downloaderx 好用记得给颗星！

### 特性

* 简单友好的图形用户界面
* 支持 Windows、Mac、Linux
* 通过 Cookie [BDUSS] 登录，无需担心密码泄漏
* 多线程下载
* 支持断点续传

![主界面](https://img.hacpai.com/file/2018/05/241876d353a447b69042a49b97d44caa_.png)

### 用法

1. [下载发布包](https://share.weiyun.com/57zViCm)
2. 在浏览器登录百度网盘，获取 BDUSS Cookie 值（**注意拷贝完整**）
   ![BDUSSpng](https://img.hacpai.com/file/2018/04/d1a78d5163f644d7931925ef5edbf9dd_BDUSS.png)
3. 登录成功后显示主界面，选择需要下载的文件后点击下载即可

![BND 教学](https://img.hacpai.com/file/2018/05/c87225b75d12411ca5ec4a57274371eb_.gif)

### 常见问题

#### 如何获取 BDUSS？

请参考[这里](https://www.baidu.com/s?wd=如何获取BDUSS)。

#### 文件列表显示不全？

只显示最新的 10 个文件（夹），可在网盘中将你要下载的文件复制到根目录，然后在 BND 中刷新就可以看到了。

#### 速度忽高忽低？

BND 显示的速度是瞬时速度，仅作为参考，不必在太在意。

#### 360 报毒？

这个我也不知道为啥，但我敢拍着胸脯保证无毒无害，业界良心！

#### Mac/Linux 下如何粘贴 BDUSS？

Mac/Linux 下在输入 BDUSS 时请使用鼠标右键的下拉菜单，不支持快捷键粘贴。

### 问题反馈

遇到问题可到[论坛](https://hacpai.com/tag/BND)发帖求助。

### 更新日志

请看[这里](https://github.com/b3log/baidu-netdisk-downloaderx/blob/master/CHANGE_LOGS.md)。

### 鸣谢

* [aria2](https://github.com/aria2/aria2)：超高速的下载引擎
* [BaiduPCS-Go](https://github.com/iikira/BaiduPCS-Go)：百度网盘客户端 - Go 语言编写
* [andlabs/ui](https://github.com/andlabs/ui)：跨平台的 Go GUI 库