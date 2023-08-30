---
title: 准备工作
lang: 'zh-CN'
permalink: /start/config
---
## 下载安卓软件（必须）
* [HttpCanary](https://ggfun.lanzoub.com/iIFIm061yhpe)
* [VMOS Pro](https://ggfun.lanzoub.com/iKUDZ061yapc)
* [Root Explorer](https://ggfun.lanzoub.com/iBFip061ykkh)

## 下载电脑端压缩包
* 由于华为云函数开始收费，所以下载包到本地不再必须，可以直接在[云端Fork](https://github.com/superggfun/smoba_LeanCloud)到自己的账户。
* [适配华为云函数](https://ggfun.lanzoub.com/iD0Gj061z9zc)

## 搭建操作环境
1. 打开HttpCanary，进入设置。选择“HttpCanary根证书”中的“导出HttpCanary根证书”，选择“System Trusted(.0)”类型导出。
2. 打开VMOS Pro，下载并添加一个虚拟机，什么版本的虚拟机无所谓，只要有Root功能。
3. 进入虚拟机之后，打开设置，勾选超级用户。
4. 打开右下角的“常用工具”，选择应用或者安装包，导入王者营地、QQ和Root Explorer。
5. 继续在“常用工具”中，选择文件，内部储存，打开路径“根目录>内部存储>HttpCanary>certs”，选择刚才导出的.0格式的证书文件，然后导入。
6. 在虚拟机中打开Root Explorer，它将会请求超级用户访问权限，选择永久记住选择，并允许。
7. 打开VMOSfiletransferstation文件夹，找到刚才导入的.0格式证书文件，长按文件复制。
8. 进入到根目录，打开路径“etc/security/cacerts”，将刚才复制的.0文件粘贴到这个位置。

## 抓包教程
1. 打开真机下的HttpCanary。点击“目标应用”，选择VMOS Pro，点击右下角蓝色的圆形按钮开始抓包。
2. 在虚拟机中登入王者营地并执行签到操作。
3. 返回真机，进入HttpCanary再次点击蓝色按钮，关闭抓包。
4. 点击HttpCanary右上角三个点里面的过滤，服务器Host。
5. 勾选ssl.kohsocialapp.qq.com域名，然后返回到HttpCanary主界面。
6. 找到“ssl.kohsocialapp.qq.com/user/login”,并点击。点击上面的“请求”，再点击下面的“Text”，找到配置config.json文件对应着的“accessToken”和“openId”的值并填入config.json。
7. 找到“kohcamp.qq.com/operation/action/signin”,并点击。点击上面的“请求”，再点击下面的“Text”，再次找到配置config.json文件中对应的“sig”、“timestamp”、“msdkEncodeParam”和“msdkToken”并填入config.json。