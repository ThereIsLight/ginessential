[TOC]
# 2020年7月13日
## 功能
实现了用户注册 + 项目重构

http的处理函数放在controller里面
route放在routes.go里面
model放在model里面
数据库的初始化放在common.database.go里面
## 问题
- go mod里面如何添加依赖：
运行时go mod自动生成依赖

数据库链接部分的代码居然一次性就写对了，牛B啊。
## PS
# 2020年7月15日
## 功能
- 用户登录
- jwt [参考链接](https://baijiahao.baidu.com/s?id=1608021814182894637&wfr=spider&for=pc)

## 下载了很多的依赖包，那些依赖包是如何管理的？
