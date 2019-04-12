# mangmang
产品目的：个人用户、公司管理、项目管理


# 启动前置条件
1. mysql
2. docker
3. docker-compose


# 部署
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
# 配置
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

# 功能
- 个人管理
- 公司管理
- 项目管理
