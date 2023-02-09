# tiktok-lite

一个简易版本的抖音server，使用gin+gorm+gen+mysql开发

接口文件在`api/`目录下

数据库连接可根据自己配置做修改，配置文件在`conf/db.go`文件

数据库表结构在`conf/dump.sql`文件

测试数据在`conf/data/`目录下

gen生成model文件在`util.mysql.go`文件

// TODO

1. 优化并发
2. 分布式重构