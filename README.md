# golang+vue 的个人博客

该博客(暂时不可以访问)部署[地址]https://(www.zerotower.com)

## 采用的技术栈

### 前端(web页面)
* vue3,前端框架
* pinia,状态管理
* vue-router, vue的路由库
* vite,打包工具
* vuetify,组件库(比element好看)
* axiso, http请求库
### 前端(管理端)

### 后端
* golang
* MySQL
* GORM
* logrus
* Redis
* ElasticSearch

## 启动项目
必须先保证后端启动
## 后端
```bash
cd server #进入到server目录
go install  # 安装golang的依赖包
go run main.go #启动服务
#------以下是表结构迁移命令选项---------#
go run main.go -db #根据表结构同步数据库
go run main.go -es #同步表结构到es
```
数据库的配置和其他项目配置请参考项目根目录下的[settings.yaml](./server/settings.yaml)。
## 前端