import axios, { AxiosResponse } from 'axios';
import { message } from 'antd';
import type {
  APIResponse,
  PaginationResponse,
  User,
  LoginRequest,
  LoginResponse,
  CreateUserRequest,
  Project,
  CreateProjectRequest,
  Pipeline,
  CreatePipelineRequest,
  Build,
  BuildStep,
  TriggerBuildRequest,
} from '@/types';

// 创建axios实例
const api = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response: AxiosResponse<APIResponse>) => {
    const { data } = response;
    if (data.code !== 200 && data.code !== 201) {
      message.error(data.message || '请求失败');
      return Promise.reject(new Error(data.message || '请求失败'));
    }
    return response;
  },
  (error) => {
    if (error.response?.status === 401) {
      message.error('登录已过期，请重新登录');
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    } else {
      message.error(error.response?.data?.message || '网络错误');
    }
    return Promise.reject(error);
  }
);

// 认证相关API
export const authAPI = {
  login: (data: LoginRequest): Promise<AxiosResponse<APIResponse<LoginResponse>>> =>
    api.post('/auth/login', data),
};

// 用户相关API
export const userAPI = {
  getProfile: (): Promise<AxiosResponse<APIResponse<User>>> =>
    api.get('/users/profile'),
  
  updateProfile: (data: Partial<User>): Promise<AxiosResponse<APIResponse<User>>> =>
    api.put('/users/profile', data),
  
  list: (page: number, pageSize: number): Promise<AxiosResponse<APIResponse<PaginationResponse<User>>>> =>
    api.get('/users', { params: { page, page_size: pageSize } }),
  
  create: (data: CreateUserRequest): Promise<AxiosResponse<APIResponse<User>>> =>
    api.post('/users', data),
  
  getById: (id: number): Promise<AxiosResponse<APIResponse<User>>> =>
    api.get(`/users/${id}`),
  
  update: (id: number, data: Partial<User>): Promise<AxiosResponse<APIResponse<User>>> =>
    api.put(`/users/${id}`, data),
  
  delete: (id: number): Promise<AxiosResponse<APIResponse>> =>
    api.delete(`/users/${id}`),
};

// 项目相关API
export const projectAPI = {
  list: (page: number, pageSize: number): Promise<AxiosResponse<APIResponse<PaginationResponse<Project>>>> =>
    api.get('/projects', { params: { page, page_size: pageSize } }),
  
  create: (data: CreateProjectRequest): Promise<AxiosResponse<APIResponse<Project>>> =>
    api.post('/projects', data),
  
  getById: (id: number): Promise<AxiosResponse<APIResponse<Project>>> =>
    api.get(`/projects/${id}`),
  
  update: (id: number, data: Partial<Project>): Promise<AxiosResponse<APIResponse<Project>>> =>
    api.put(`/projects/${id}`, data),
  
  delete: (id: number): Promise<AxiosResponse<APIResponse>> =>
    api.delete(`/projects/${id}`),
  
  getMyProjects: (): Promise<AxiosResponse<APIResponse<Project[]>>> =>
    api.get('/projects/my'),
};

// 流水线相关API
export const pipelineAPI = {
  list: (page: number, pageSize: number): Promise<AxiosResponse<APIResponse<PaginationResponse<Pipeline>>>> =>
    api.get('/pipelines', { params: { page, page_size: pageSize } }),
  
  create: (data: CreatePipelineRequest): Promise<AxiosResponse<APIResponse<Pipeline>>> =>
    api.post('/pipelines', data),
  
  getById: (id: number): Promise<AxiosResponse<APIResponse<Pipeline>>> =>
    api.get(`/pipelines/${id}`),
  
  update: (id: number, data: Partial<Pipeline>): Promise<AxiosResponse<APIResponse<Pipeline>>> =>
    api.put(`/pipelines/${id}`, data),
  
  delete: (id: number): Promise<AxiosResponse<APIResponse>> =>
    api.delete(`/pipelines/${id}`),
  
  getByProject: (projectId: number): Promise<AxiosResponse<APIResponse<Pipeline[]>>> =>
    api.get(`/pipelines/project/${projectId}`),
};

// 构建相关API
export const buildAPI = {
  list: (page: number, pageSize: number): Promise<AxiosResponse<APIResponse<PaginationResponse<Build>>>> =>
    api.get('/builds', { params: { page, page_size: pageSize } }),
  
  create: (data: TriggerBuildRequest): Promise<AxiosResponse<APIResponse<Build>>> =>
    api.post('/builds', data),
  
  getById: (id: number): Promise<AxiosResponse<APIResponse<Build>>> =>
    api.get(`/builds/${id}`),
  
  updateStatus: (id: number, status: string): Promise<AxiosResponse<APIResponse>> =>
    api.put(`/builds/${id}/status`, { status }),
  
  getSteps: (id: number): Promise<AxiosResponse<APIResponse<BuildStep[]>>> =>
    api.get(`/builds/${id}/steps`),
  
  getByPipeline: (pipelineId: number, page: number, pageSize: number): Promise<AxiosResponse<APIResponse<PaginationResponse<Build>>>> =>
    api.get(`/builds/pipeline/${pipelineId}`, { params: { page, page_size: pageSize } }),
};

export default api; 