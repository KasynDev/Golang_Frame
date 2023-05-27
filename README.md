# Golang_Frame
一个Go Web开发的框架，使用gin进行搭建，完成了结构化参数配置，日志打印，优雅关机和平滑重启等设置，更改config.yaml配置后可开箱即用



## 框架简介

1. 可作为go gin开发框架
2. 使用sqlx连接mysql（可更换为gorm）
3. 利用zap日志库建立日志信息输出
4. 使用viper库加载配置信息，并使用结构体变量优化
5. 实现优雅关机和平滑重启



## 文件目录

```bash
Golang_Frame
├─conf           // 存放配置信息
├─controllers    // 主要操作代码（可选）
├─dao            // 数据库对象
│  ├─mysql       // 配置mysql连接
│  └─redis       // 配置redis连接
├─logger         // 配置日志
├─logic          // 放置逻辑代码（可选）
├─models         // 放置模块化代码（可选）
├─pkg            // 放置第三方包（可选）
├─routes         // 配置路由
└─settings       // 映射配置信息
```



## 项目启动

项目启动非常简单，只需要正确配置好config.yaml文件内容即可，以下为一个简单示例：

```yaml
app:
  name: "go_frame"               // Web-app名称
  mode: "dev"                    // 模式
  version: "1.0.0"               // 版本号
  port: ":8081"                  // 端口号

log:                             // 日志信息
  level: "debug"                 // 日志等级（debug, info, error）
  filename: "go_frame.log"       // 输出的日志文件名称
  max_size: 200                  // 日志最大容量（单位为MB）
  max_age: 30                    // 日志最长储存天数
  max_backups: 7                 // 日志备份数量

mysql:                           // mysql数据库
  host: "127.0.0.1"              // Host主机号
  port: 3306                     // mysql数据库端口号（default: 3306）
  user: "root"                   // mysql数据库用户名
  password: "root"               // mysql数据库密码
  db_name: "sql_demo"            // 选择想要连接的数据库
  max_open_conns: 200            // 最大连接数
  max_idle_conns: 50             // 空闲连接的最大数量

redis:                           // redis信息
  host: "127.0.0.1"              // Host主机号
  port: 6379                     // redis端口号（defalut: 6379）
  password: ""                   // redis连接密码
  db: 0                          // 选择数据库
  pool_size: 100                 // 连接池大小
```



配置完成之后，使用

```go
go mod tidy
```

安装所有所需的第三方库，完成配置后即可正常启动程序。
