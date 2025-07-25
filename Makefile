.PHONY: help setup dev build clean test

# 默认目标
help:
	@echo "Vortexia 项目构建工具"
	@echo ""
	@echo "可用命令:"
	@echo "  setup     - 初始化项目依赖"
	@echo "  dev       - 启动开发环境"
	@echo "  build     - 构建项目"
	@echo "  test      - 运行测试"
	@echo "  clean     - 清理构建产物"
	@echo "  docker    - 构建Docker镜像"
	@echo "  up        - 启动所有服务"
	@echo "  down      - 停止所有服务"

# 初始化项目
setup:
	@echo "🚀 初始化项目..."
	chmod +x scripts/setup.sh
	./scripts/setup.sh

# 开发环境
dev:
	@echo "🔧 启动开发环境..."
	docker-compose up -d postgres redis
	@echo "✅ 数据库服务已启动"
	@echo "请在不同终端运行:"
	@echo "  cd backend && go run cmd/server/main.go"
	@echo "  cd frontend && npm run dev"

# 构建项目
build:
	@echo "🏗️ 构建项目..."
	cd backend && go build -o ../bin/server cmd/server/main.go
	cd frontend && npm run build
	@echo "✅ 构建完成"

# 运行测试
test:
	@echo "🧪 运行测试..."
	cd backend && go test ./...
	cd frontend && npm run test
	@echo "✅ 测试完成"

# 清理
clean:
	@echo "🧹 清理构建产物..."
	rm -rf bin/
	rm -rf frontend/dist/
	docker system prune -f
	@echo "✅ 清理完成"

# 构建Docker镜像
docker:
	@echo "🐳 构建Docker镜像..."
	docker build -t vortexia-backend:latest backend/
	cd frontend && npm run build
	docker build -t vortexia-frontend:latest -f docker/frontend.Dockerfile .
	@echo "✅ Docker镜像构建完成"

# 启动所有服务
up:
	@echo "🚀 启动所有服务..."
	docker-compose up -d
	@echo "✅ 所有服务已启动"

# 停止所有服务
down:
	@echo "🛑 停止所有服务..."
	docker-compose down
	@echo "✅ 所有服务已停止"

# 查看日志
logs:
	docker-compose logs -f

# 数据库迁移
migrate:
	@echo "📊 运行数据库迁移..."
	cd backend && go run github.com/pressly/goose/v3/cmd/goose postgres "host=localhost port=5432 user=ci_user password=ci_password dbname=vortexia_db sslmode=disable" up
	@echo "✅ 数据库迁移完成" 