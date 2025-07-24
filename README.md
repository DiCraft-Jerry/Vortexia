# Simple CI/CD

一个基于 Go + React 构建的现代化 CI/CD 流水线工具

## 🚀 技术栈

### 后端

- **语言**: Go 1.21+
- **框架**: Gin (HTTP路由)
- **数据库**: PostgreSQL + Redis
- **认证**: JWT
- **API文档**: Swagger
- **测试**: Testify

### 前端

- **语言**: TypeScript
- **框架**: React 18
- **构建工具**: Vite
- **UI库**: Ant Design
- **状态管理**: Zustand
- **HTTP客户端**: Axios

### 基础设施

- **容器化**: Docker + Docker Compose
- **数据库迁移**: Goose
- **监控**: Prometheus + Grafana
- **日志**: Structured logging with zap

## 📁 项目结构

```
simple-ci/
├── backend/                 # Go后端服务
│   ├── cmd/                # 应用入口
│   ├── internal/           # 内部包
│   │   ├── api/           # API处理器
│   │   ├── service/       # 业务逻辑
│   │   ├── repository/    # 数据访问层
│   │   ├── model/         # 数据模型
│   │   └── middleware/    # 中间件
│   ├── pkg/               # 可复用包
│   ├── configs/           # 配置文件
│   ├── migrations/        # 数据库迁移
│   └── docs/             # API文档
├── frontend/              # React前端应用
│   ├── src/
│   │   ├── components/   # React组件
│   │   ├── pages/        # 页面组件
│   │   ├── stores/       # 状态管理
│   │   ├── services/     # API服务
│   │   ├── types/        # TypeScript类型
│   │   └── utils/        # 工具函数
│   ├── public/           # 静态资源
│   └── dist/             # 构建输出
├── docker/               # Docker配置
├── scripts/              # 脚本工具
├── docs/                 # 项目文档
└── deployments/          # 部署配置
```

## 🛠️ 快速开始

### 前置要求

- Go 1.21+
- Node.js 18+
- Docker & Docker Compose
- PostgreSQL
- Redis

### 开发环境启动

1. **克隆项目**

```bash
git clone <repository-url>
cd simple-ci
```

2. **启动基础服务**

```bash
docker-compose up -d postgres redis
```

3. **启动后端服务**

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

4. **启动前端服务**

```bash
cd frontend
npm install
npm run dev
```

### 访问地址

- 前端应用: http://localhost:3000
- 后端API: http://localhost:8080
- API文档: http://localhost:8080/swagger/index.html

## 📚 开发指南

详细的开发指南请参考 [docs/development.md](docs/development.md)

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情
