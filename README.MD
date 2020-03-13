# Log kig By Go

作为一个中间统计平台，可以快速分析API频率、可用性等；日志储存，使用唯一标识储存（如订单号）、指定查询。

## 使用前提

基于go语言、goframe框架。「base on Go Language，GoFrame frame work」

## 安装步骤

先下载源码构建 「Download and build」

保留config、public目录部署，修改config/config.toml文件中的数据库配置「edit config.toml file to save your database configure」

启动，访问`127.0.0.1:8199(默认情况)`  「run and access `127.0.0.1:8199` in browser」

你可以看到这个页面，开始安装「import mysql struct」

[![8udJE9.png](https://s1.ax1x.com/2020/03/13/8udJE9.png)](https://imgchr.com/i/8udJE9)

安装完成后添加项目 「add the project after install」

查看上报格式说明教程 上报你的数据 「learn how to report data, and do it !」

[Report上报数据格式说明](./document/reportData.md)

[php上报数据demo](./document/reportCode/php.md)

在平台中搜索查看吧! 「can search data when you need」

遇到问题希望能反馈给我改进! 「if it have question ,please tell me」

## 使用场景

- 日志查询 一般用于自己记录程序的运行状态使用。使用订单号做记录，如果某个订单出问题，可以查询出整条调用链
- api统计  用于接口可用性、性能分析，可以查到平均耗时、失败次数等，在SAAS等类型平台时还可以加上第三方客户的订单号，可以协助客户分析接口接入，加快对接

## 截图一览

![8u6EY4.png](https://s1.ax1x.com/2020/03/13/8u6EY4.png)

--------

![8u6Wn0.png](https://s1.ax1x.com/2020/03/13/8u6Wn0.png)