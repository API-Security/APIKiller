<<<<<<< HEAD
# APIKiller
API Security DAST &amp; Operations
=======
# APIKIller


## 项目简介
**一款高度可定制化的DAST&API Security平台**

详细参考：

## Feature
- 支持HTTP/HTTPS流量检测
- 多来源检测
  - 支持流量监听
  - 支持历史流量回扫\[目前只支持burpsuite保存流量\]
- 多功能扫描模块
  - 越权检测模块，高效精准，支持多情景检测
    - 具备多角色账号、单角色账号测试能力【单角色账号测试暂不启用，黑盒师傅有想法的可以直接看文档中的介绍，或私下交流】
    - 多维度、特征化判断引擎
  - csrf检测模块
    - 支持token检测
    - 常见的referer、origin检测
  - 【欢迎大家积极提PR】
- 多功能Filter处理，默认自带多个filter
  - 针对性扫描，例如只对 baidu.com域名进行扫描
  - 去重扫描，提高效率
  - 自动过滤静态文件(js,gif,jpg,png,css,jpeg,xml,img,svg...)
- API 运维
  - 提供简易的API Security运维平台
- 多方式漏洞发现提醒
  - Lark飞书
  - ...
- 对抗常见风控手段
  - 频控
- **【重磅】以上都可以快速进行拓展&二次开发**

## 食用宝典
1. 安装好数据库环境（我个人采用的是docker）
   1. docker pull 数据库镜像
   ```shell
    sudo docker run --name mysql-server -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306  mysql:5.7
   ```
   2. 导入apikiller.sql文件
   ```shell
    sudo docker cp /tmp/apikiller.sql bteye-mysql:/tmp/apikiller.sql
   ```
   3. 登入mysql
   ```shell
    docker exec -it mysql-server mysql -uroot -p123456
    source /tmp/apikiller.sql
   ```
   4. 在 config.json 中配置好
   
   ![img_6.png](static/img/img_6.png)
2. 安装根目录下的https证书[windows环境]
   1. 找到根目录下的ca.crt证书  
   ![img.png](static/img/img.png)
   2. 点击安装即可，将其添加到系统根信任证书颁发位置
3. 配置漏洞发现通知Webhook
   1. 根据[飞书指导](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN)，开启一个bot，并复制相关的webhook【支持secret鉴权操作】
   2. 在根路径下的config.json中进行配置(如果有secret，就进行配置)
   ![img_1.png](static/img/img_1.png)
   3. 配置成功后，当发现漏洞时，会立即推送漏洞信息
   ![img_2.png](static/img/img_2.png)
4. 一键启动【所有的配置都在config.json】

    ![img_4.png](static/img/img_4.png)
    ![img_3.png](static/img/img_3.png)
5. ding~，发现新漏洞，快去看鸭

## 检测效果展示
这里基于[VAPI越权靶场](https://www.freebuf.com/vuls/332312.html) 进行实战模拟
配好环境后，先根据项目鉴权机制，拿两个账号，配置好config.json
![img.png](static/img/img_20.png)
然后启动项目，访问接口
![img_1.png](static/img/img_21.png)
![img.png](static/img/img_22.png)
**成功检测出越权和csrf**
![img_3.png](static/img/img_23.png)

## 项目架构
![architecture](static/img/architecture.jpg)


## API运营平台
![](static/img/2.jpg)

## 二次开发文档
【暂无】

## 优化记录


## 后期计划
- [ ] 其他数据库支持
- [ ] 添加基于poc的扫描模块
- [ ] 403bypass module
- [ ] 完善其他通知方式的支持
- [ ] 支持扫描代理功能
- [ ] 提供docker版快速搭建能力
- [ ] 其他

## 感谢
【**最后感谢项目中所使用到的各种开源组件的作者**】

>>>>>>> ee7bd79 (v0.0.1)
