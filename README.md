# mangmang
产品功能：个人管理、公司管理、项目管理


## 启动前置条件
1. mysql5
2. docker
3. docker-compose


## 部署
1. 拉取仓库
    ```
    git https://github.com/qzq1111/mangmang.git
    ```
2. 构建镜像
    ```
    cd mangmang
    docker build . -t=mangmang:latest
    ```
3. 运行容器
    ```
    docker-compose up -d
    ```
4. 访问
    ```
    http://127.0.0.1:3010
    ```
## 配置
1. 修改配置文件 conf/app.ini
    ```
    # app 通用配置
    [app]
    Page = 1
    PageSize = 10
    JwtSecret = 23347$040412 
    AvatarPath = ./AvatarFile # 头像上传位置
    
    [server]
    RunMode = debug # 启动模式 debug 或者release
    HTTPPort = 80  # 启动端口
    ReadTimeout = 60 # 读取超时
    WriteTimeout = 60 # 写入超时
    
    # 数据库配置
    [database]
    Type = mysql # 数据库连接属性
    User = root # 数据库用户
    Password = mad123 # 用户密码
    Host = 127.0.0.1:3306 # 数据库地址
    Name = mangmang  # 数据库名称
    
    # redis配置
    [redis]
    Host = 127.0.0.1:6379
    Password =
    MaxIdle = 30
    MaxActive = 30
    IdleTimeout = 200
    ```
2. 重新build Docker镜像
    ```
    docker build . -t=mangmang:latest
    ```
3. 启动
    ```
    docker-compose up -d
    ```

## 功能（待完善）
- [ ] 个人管理
    - [x] 登陆注册
    - [x] 个人信息设置
    - [x] 个人修改密码
    - [x] 个人名片设置
    - [ ] 好友添加删除
    - [ ] 好友聊天
    
- [ ] 公司管理
    - [ ] 公司信息
    - [ ] 部门管理
    - [ ] 考勤
    - [ ] 报销
    
- [ ] 项目管理
    - [x] 项目新建修改
    - [x] 项目成员管理
    - [x] 成员角色设置
    - [ ] 项目任务分发
    - [ ] 项目统计信息
    - [ ] 个人任务管理
    
## 感谢
感谢[@EDDYCJY](https://github.com/EDDYCJY)的文章[go-gin-example](https://github.com/EDDYCJY/go-gin-example).
