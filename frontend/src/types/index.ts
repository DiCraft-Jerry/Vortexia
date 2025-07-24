// API响应类型
export interface APIResponse<T = any> {
  code: number;
  message: string;
  data?: T;
}

// 分页响应类型
export interface PaginationResponse<T = any> {
  items: T[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

// 用户相关类型
export interface User {
  id: number;
  username: string;
  email: string;
  role: 'admin' | 'user';
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface CreateUserRequest {
  username: string;
  email: string;
  password: string;
  role: 'admin' | 'user';
}

// 项目相关类型
export interface Project {
  id: number;
  name: string;
  description: string;
  repo_url: string;
  branch: string;
  owner_id: number;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateProjectRequest {
  name: string;
  description: string;
  repo_url: string;
  branch: string;
}

// 流水线相关类型
export interface Pipeline {
  id: number;
  project_id: number;
  name: string;
  config: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreatePipelineRequest {
  project_id: number;
  name: string;
  config: string;
}

// 构建相关类型
export interface Build {
  id: number;
  pipeline_id: number;
  branch: string;
  commit: string;
  status: 'pending' | 'running' | 'success' | 'failed' | 'canceled';
  started_at: string;
  finished_at?: string;
  duration?: number;
  trigger_by: number;
  created_at: string;
}

export interface BuildStep {
  id: number;
  build_id: number;
  name: string;
  command: string;
  status: 'pending' | 'running' | 'success' | 'failed' | 'skipped';
  output: string;
  started_at: string;
  finished_at?: string;
  duration?: number;
  step_order: number;
}

export interface TriggerBuildRequest {
  pipeline_id: number;
  branch: string;
  commit?: string;
}

// 状态相关类型
export type BuildStatus = 'pending' | 'running' | 'success' | 'failed' | 'canceled';
export type StepStatus = 'pending' | 'running' | 'success' | 'failed' | 'skipped'; 