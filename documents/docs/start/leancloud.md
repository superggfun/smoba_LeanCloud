---
title: 部署在LeanCloud（推荐）
lang: 'zh-CN'
permalink: /start/leancloud
---

## 创建云函数
1. 请在GitHub上先[分叉（Fork）LeanCloud仓库](https://github.com/superggfun/smoba_LeanCloud)到您的账户，并确保其权限为公开（Public）。
2. 接下来，在Github的相应目录中编辑config.json文件。详细内容请[参考此处](/start/config/#抓包教程)。
3. 接着，[登录LeanCloud官方网站](https://www.leancloud.cn/)或选择[国际版登录](https://www.leancloud.app/)。
4. 操作路径如下：登录到LeanCloud > 选择“创建应用” > 进入“开发板” > 选择“云引擎” > 点击“管理部署” > 在“分组”中选择“创建分组” > 进入“云函数和Hooks” > 选择“实例规格”为“免费版” > 选择“git部署” > 在“配置git”处将您分叉后的仓库地址粘贴进去 > 点击“保存” > 选择“部署”。部署完成后稍作刷新，在“云函数”区域找到“run”并点击运行。


<iframeComp ihtml="https://liveimages.videocc.net/uploadimage/20221011/aliyun?u=https://s3plus.meituan.net/v1/mss_550586ef375b493da4aa79bebdfce4fa/csc-apply-file-web/prod/2023-08-29/4178199b-b7c0-45dc-8c0d-d8fd4e6506dc.m3u8"></iframeComp>

## 定时触发
1. 导航至云引擎 > 管理部署 > 所创建的分组 > 定时任务，然后点击“创建定时任务”。
2. 任务名称可根据需求任意命名。在“运行目标”中选择“云函数”，并设置为“生产环境”下的“run”。
2. 对于“运行周期”，请使用Cron表达式。建议使用“0 0 20 * * ?”来于每晚8点执行，或“0 0 23 * * ?”于每晚11点执行

![alt text](/leancloud.jpg)