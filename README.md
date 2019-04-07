# mangmang
实现公司管理、个人用户项目管理、项目管理

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

# 功能
- 个人管理
- 公司管理
- 项目管理
