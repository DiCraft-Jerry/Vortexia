#!/bin/bash

# Simple CI/CD 项目设置脚本

echo "🚀 正在设置 Simple CI/CD 项目..."

# 检查必要的工具
check_tool() {
    if ! command -v $1 &> /dev/null; then
        echo "❌ $1 未安装，请先安装 $1"
        exit 1
    else
        echo "✅ $1 已安装"
    fi
}

echo "📋 检查依赖工具..."
check_tool "go"
check_tool "node"
check_tool "docker"
check_tool "docker-compose"

# 设置后端
echo "🔧 设置 Go 后端..."
cd backend
go mod tidy
go mod download

# 设置前端
echo "🎨 设置 React 前端..."
cd ../frontend
npm install

# 返回根目录
cd ..

# 复制环境变量文件
echo "📝 设置环境变量..."
if [ ! -f backend/.env ]; then
    cp backend/.env.example backend/.env
    echo "✅ 已创建 backend/.env 文件，请根据需要修改配置"
fi

echo "🎉 项目设置完成！"
echo ""
echo "📖 使用说明："
echo "1. 启动数据库服务: docker-compose up -d postgres redis"
echo "2. 运行数据库迁移: cd backend && go run migrations/*.sql"
echo "3. 启动后端服务: cd backend && go run cmd/server/main.go"
echo "4. 启动前端服务: cd frontend && npm run dev"
echo ""
echo "🌐 访问地址："
echo "- 前端应用: http://localhost:3000"
echo "- 后端API: http://localhost:8080"
echo "- API文档: http://localhost:8080/swagger/index.html"
echo ""
echo "👤 默认管理员账号："
echo "- 用户名: admin"
echo "- 密码: admin123" 