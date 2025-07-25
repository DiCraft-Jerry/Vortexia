-- +goose Up
-- 创建用户表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 创建项目表
CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    repo_url VARCHAR(500) NOT NULL,
    branch VARCHAR(100) NOT NULL DEFAULT 'main',
    owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 创建流水线表
CREATE TABLE pipelines (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    config TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 创建构建表
CREATE TABLE builds (
    id SERIAL PRIMARY KEY,
    pipeline_id INTEGER NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
    branch VARCHAR(100) NOT NULL,
    commit VARCHAR(40),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMP WITH TIME ZONE,
    duration INTEGER, -- 构建持续时间（秒）
    trigger_by INTEGER NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 创建构建步骤表
CREATE TABLE build_steps (
    id SERIAL PRIMARY KEY,
    build_id INTEGER NOT NULL REFERENCES builds(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    command TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    output TEXT,
    started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMP WITH TIME ZONE,
    duration INTEGER, -- 步骤持续时间（秒）
    step_order INTEGER NOT NULL DEFAULT 0
);

-- 创建索引
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_active ON users(is_active);

CREATE INDEX idx_projects_owner ON projects(owner_id);
CREATE INDEX idx_projects_active ON projects(is_active);

CREATE INDEX idx_pipelines_project ON pipelines(project_id);
CREATE INDEX idx_pipelines_active ON pipelines(is_active);

CREATE INDEX idx_builds_pipeline ON builds(pipeline_id);
CREATE INDEX idx_builds_status ON builds(status);
CREATE INDEX idx_builds_created ON builds(created_at);

CREATE INDEX idx_build_steps_build ON build_steps(build_id);
CREATE INDEX idx_build_steps_status ON build_steps(status);
CREATE INDEX idx_build_steps_order ON build_steps(step_order);

-- 插入默认管理员用户
-- 密码是 admin123 的哈希值
INSERT INTO users (username, email, password_hash, role, is_active) 
VALUES ('admin', 'admin@vortexia.com', '$2a$10$D4L/4.e/LZJ8MdZ8l9U./.WuGF4ByX4vP7vQ3KQXKXVWxqWYpSd1C', 'admin', true);

-- +goose Down
DROP TABLE IF EXISTS build_steps;
DROP TABLE IF EXISTS builds;
DROP TABLE IF EXISTS pipelines;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users; 