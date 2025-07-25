# Vortexia 开发指南

## 📋 项目概述

Vortexia 是一个基于 Go + React 构建的轻量级持续集成/持续部署平台，针对 1核2GB 服务器环境进行了优化。

## 🛠️ 技术栈

### 后端

- **语言**: Go 1.21+
- **框架**: Gin
- **数据库**: PostgreSQL + Redis
- **认证**: JWT
- **文档**: Swagger
- **日志**: Zap

### 前端

- **语言**: TypeScript
- **框架**: React 18
- **构建工具**: Vite
- **UI库**: Ant Design
- **状态管理**: Zustand
- **HTTP**: Axios

## 🚀 快速开始

### 1. 环境准备

```bash
# 检查工具版本
go version    # >= 1.21
node -v       # >= 18
docker -v     # 最新版本
```

### 2. 项目初始化

```bash
# 使用自动化脚本
make setup

# 或手动设置
git clone <项目地址>
cd Vortexia

# 后端依赖
cd backend && go mod tidy

# 前端依赖  
cd ../frontend && npm install
```

### 3. 启动开发环境

```bash
# 启动数据库服务
make dev

# 在不同终端中启动服务
# 终端1: 后端服务
cd backend && go run cmd/server/main.go

# 终端2: 前端服务
cd frontend && npm run dev
```

### 4. 访问应用

- 前端: http://localhost:3000
- 后端API: http://localhost:8080
- API文档: http://localhost:8080/swagger/index.html

**默认管理员账号:**

- 用户名: `admin`
- 密码: `admin123`

## 📁 项目结构详解

### 后端结构

```
backend/
├── cmd/server/         # 应用入口
├── internal/
│   ├── api/
│   │   ├── handlers/   # HTTP处理器
│   │   └── routes/     # 路由配置
│   ├── service/        # 业务逻辑层
│   ├── repository/     # 数据访问层
│   ├── model/          # 数据模型
│   ├── middleware/     # 中间件
│   └── config/         # 配置管理
├── pkg/                # 可复用包
├── migrations/         # 数据库迁移
└── docs/              # API文档
```

### 前端结构

```
frontend/
├── src/
│   ├── components/     # React组件
│   │   ├── common/     # 通用组件
│   │   ├── layout/     # 布局组件
│   │   └── forms/      # 表单组件
│   ├── pages/          # 页面组件
│   ├── stores/         # 状态管理
│   ├── services/       # API服务
│   ├── types/          # TypeScript类型
│   ├── utils/          # 工具函数
│   └── assets/         # 静态资源
└── public/             # 公共文件
```

## 💻 开发流程

### 1. 添加新功能

#### 后端开发流程

```bash
# 1. 定义数据模型 (internal/model/)
# 2. 创建数据库表 (migrations/)
# 3. 实现Repository接口 (internal/repository/)
# 4. 实现Service业务逻辑 (internal/service/)
# 5. 创建API处理器 (internal/api/handlers/)
# 6. 配置路由 (internal/api/routes/)
```

#### 前端开发流程

```bash
# 1. 定义TypeScript类型 (src/types/)
# 2. 创建API服务 (src/services/)
# 3. 实现页面组件 (src/pages/)
# 4. 添加路由配置 (src/App.tsx)
# 5. 更新状态管理 (src/stores/)
```

### 2. 数据库操作

```bash
# 创建新迁移
cd backend
go run github.com/pressly/goose/v3/cmd/goose create add_new_table sql

# 运行迁移
make migrate

# 回滚迁移
cd backend && goose postgres "connection_string" down
```

### 3. API文档

```bash
# 生成Swagger文档
cd backend
swag init -g cmd/server/main.go

# 访问文档
http://localhost:8080/swagger/index.html
```

## 🧪 测试

### 后端测试

```bash
# 运行所有测试
cd backend && go test ./...

# 运行特定包测试
go test ./internal/service/

# 生成测试覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 前端测试

```bash
# 运行测试
cd frontend && npm test

# 生成覆盖率报告
npm run test:coverage
```

## 🐳 容器化部署

### 1. 构建镜像

```bash
# 构建所有镜像
make docker

# 单独构建
docker build -t vortexia-backend backend/
docker build -t vortexia-frontend frontend/
```

### 2. 运行容器

```bash
# 使用docker-compose
make up

# 停止服务
make down

# 查看日志
make logs
```

## 🔧 配置说明

### 环境变量 (backend/.env)

```bash
# 服务器配置
SERVER_PORT=8080
GIN_MODE=release

# 数据库配置
DB_HOST=postgres
DB_PORT=5432
DB_USER=ci_user
DB_PASSWORD=ci_password
DB_NAME=vortexia_db

# Redis配置
REDIS_HOST=redis
REDIS_PORT=6379

# JWT配置
JWT_SECRET=your-secret-key
JWT_EXPIRE=7200
```

### 性能优化配置

#### PostgreSQL (configs/postgres-small.conf)

```sql
shared_buffers = 128MB
work_mem = 2MB
max_connections = 20
```

#### Go应用优化

```go
// 内存优化
debug.SetGCPercent(20)
runtime.GOMAXPROCS(1)

// 数据库连接池
db.SetMaxOpenConns(5)
db.SetMaxIdleConns(2)
```

## 🚨 故障排查

### 常见问题

1. **数据库连接失败**

   ```bash
   # 检查数据库状态
   docker ps | grep postgres

   # 查看日志
   docker logs vortexia-postgres
   ```
2. **前端依赖错误**

   ```bash
   # 清理并重装依赖
   cd frontend
   rm -rf node_modules package-lock.json
   npm install
   ```
3. **Go模块问题**

   ```bash
   # 清理模块缓存
   cd backend
   go clean -modcache
   go mod tidy
   ```

### 性能监控

```bash
# 内存使用
docker stats

# 数据库性能
docker exec -it vortexia-postgres psql -U ci_user -d vortexia_db -c "SELECT * FROM pg_stat_activity;"

# API响应时间
curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/api/v1/health"
```

## 🔐 安全最佳实践

1. **JWT密钥**: 生产环境使用强随机密钥
2. **数据库密码**: 使用复杂密码
3. **HTTPS**: 生产环境启用SSL/TLS
4. **CORS**: 配置适当的跨域策略
5. **输入验证**: 所有API输入进行验证

## 🤝 贡献指南

1. Fork项目
2. 创建特性分支: `git checkout -b feature/new-feature`
3. 提交更改: `git commit -am 'Add new feature'`
4. 推送分支: `git push origin feature/new-feature`
5. 创建Pull Request

## 📞 支持

如有问题，请通过以下方式联系：

- 创建Issue
- 查看FAQ文档
- 参考示例代码

---

**祝您开发愉快！** 🎉
