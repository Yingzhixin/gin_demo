# gin_demo

#### 介绍
从零开始搭建Go的GIN项目并部署到Docker

#### 技术选型
| 名称  | 地址                                          | 备注             |
| ----- | --------------------------------------------- | ---------------- |
| GIN   | https://github.com/gin-gonic/gin              | Web框架          |
| Gorm  | https://github.com/jinzhu/gorm                | ORM框架          |
| Swag  | https://github.com/swaggo/swag/blob/master/README_zh-CN.md | 接口文档         |
| Redis | https://github.com/go-redis/redis             | 缓存             |
| Mysql | https://github.com/jinzhu/gorm/dialects/mysql | 关系型数据库驱动   |

#### 部署
- 将项目整个文件夹拷贝到服务器上
- 导航至项目文件夹
- 使用命令``docker build -f Dockerfile -t <tag> .``根据Dockerfile创建镜像
- 使用命令``docker run -idt -p 9000:9000 <image_name>``构建容器

