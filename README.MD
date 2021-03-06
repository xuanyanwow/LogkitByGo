# Log kit By Go

作为一个中间统计平台，可以快速分析API频率、可用性等；日志储存，使用唯一标识储存（如订单号）、指定查询。

## 作者

- QQ 59419979
- Email 59419979@qq.com
- 交流群 人数较少，暂无，可以先加QQ交流，后续拉群

## 使用前提

基于go语言、goframe框架。「base on Go Language，GoFrame frame work」

运行go程序后默认开启两个监听服务

- http服务 用于后台查询展示页面 默认8199端口
- tcp服务  用于接收上报数据 默认8999端口

## 安装步骤

先下载源码构建 「Download and build」

不想构建和学习go语言，直接下载我们构建好的包吧（支持windows和linux 64位系统）

- windows [马上就有]
- linux [马上就有]

保留config、public、template目录部署，修改config/config.toml文件中的数据库配置「edit config.toml file to save your database configure」

启动，访问`127.0.0.1:8199(默认情况)`  「run and access `127.0.0.1:8199` in browser」

你可以看到这个页面，开始安装「import mysql struct」

[![8udJE9.png](https://s1.ax1x.com/2020/03/13/8udJE9.png)](https://imgchr.com/i/8udJE9)

导入数据库成功后，修改后台页面的接口地址为服务器ip+端口（默认为127.0.0.1:8199） 地址在`public/lib/layui/lay/okmodules/siamConfig.js`

完成后添加项目 「add the project after install」

查看上报格式说明教程 上报你的数据 「learn how to report data, and do it !」

[Report上报数据格式说明](./document/reportData.md)

[php上报数据demo](./document/reportCode/php.md)

[java上报数据demo](./document/reportCode/java.md)

在平台中搜索查看吧! 「can search data when you need」

遇到问题希望能反馈给我改进! 「if it have question ,please tell me」

## 使用场景

- 日志查询 一般用于自己记录程序的运行状态使用。使用订单号做记录，如果某个订单出问题，可以查询出整条调用链
- api统计  用于接口可用性、性能分析，可以查到平均耗时、失败次数等，在SAAS等类型平台时还可以加上第三方客户的订单号，可以协助客户分析接口接入，加快对接

## 截图一览

[![8u6EY4.png](https://s1.ax1x.com/2020/03/13/8u6EY4.png)](https://imgchr.com/i/8u6EY4)

--------

[![8u6Wn0.png](https://s1.ax1x.com/2020/03/13/8u6Wn0.png)](https://imgchr.com/i/8u6Wn0)
