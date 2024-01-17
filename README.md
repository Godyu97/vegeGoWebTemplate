## 项目名称
> vegeGoWebTemplate

* 参考文献:
    * https://github.com/golang-standards/project-layout/blob/master/README_zh.md
    * https://makeoptim.com/golang/standards/project-layout
    * https://go-zero.dev/

* etc:静态配置文件目录
* cmd:gorm gen的生成脚本等生成工具
* pkg:外部应用程序可以使用的库代码
* main.go:程序启动入口文件
* internal:单个服务内部文件，其可见范围仅限当前服务
    * config:静态配置文件对应的结构体声明目录
    * handler:handler 目录，可选，一般 http 服务会有这一层做路由管理，handler 为固定后缀
    * logic:业务目录，所有业务编码文件都存放在这个目录下面，logic 为固定后缀
    * svc:依赖注入目录，所有 logic 层需要用到的依赖都要在这里进行显式注入
    * types:结构体存放目录
    * dao:db操作api
    * dal:gorm gen生成的model和query