# tiktok-lite

一个简易版本的抖音server，使用gin+gorm+gen+mysql开发

接口文件在`api/`目录下

数据库连接可根据自己配置做修改，配置文件在`conf/db.go`文件

数据库表结构在`conf/dump.sql`文件

测试数据在`conf/data/`目录下

gen生成model文件在`util.mysql.go`文件


## 重构ing
`go version: 1.21`

### 微服务重构day1

因为gorm+gen太麻烦，不好用，因此考虑使用python实现微服务层

使用peewee作为orm

规划好微服务项目目录

    ├── src: 项目源代码
    │   ├── config: 配置文件
    │   ├── extra: 外部服务依赖
    │   ├── idl: proto文件
    │   ├── model: 数据模型
    │   ├── rpc: Rpc代码
    │   ├── service: 微服务实例
    │   ├── storage: 存储相关
    │   ├── util: 辅助代码
    │   └── web: API网关代码

数据库表重构完成

定义user.proto

day2 TODO:

因为config/strings中的常量在py和go中都需要用到，所以要将config/strings中的文件变成yaml文件供py和go读取
