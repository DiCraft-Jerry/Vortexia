import { useState } from 'react';
import { Form, Input, Button, Card, Typography, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { useAuthStore } from '@/stores/authStore';
import type { LoginRequest } from '@/types';

const { Title, Text } = Typography;

export default function Login() {
  const [form] = Form.useForm();
  const { login, loading } = useAuthStore();

  const handleSubmit = async (values: LoginRequest) => {
    try {
      await login(values);
      message.success('登录成功');
    } catch (error) {
      // 错误已在拦截器中处理
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <Card className="w-full max-w-md">
        <div className="text-center mb-8">
          <Title level={2} className="text-blue-600">
            Simple CI/CD
          </Title>
          <Text type="secondary">持续集成/持续部署平台</Text>
        </div>

        <Form
          form={form}
          name="login"
          onFinish={handleSubmit}
          autoComplete="off"
          size="large"
        >
          <Form.Item
            name="username"
            rules={[
              { required: true, message: '请输入用户名' },
              { min: 3, message: '用户名至少3个字符' },
            ]}
          >
            <Input
              prefix={<UserOutlined />}
              placeholder="用户名"
            />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[
              { required: true, message: '请输入密码' },
              { min: 6, message: '密码至少6个字符' },
            ]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="密码"
            />
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              className="w-full"
              loading={loading}
            >
              登录
            </Button>
          </Form.Item>
        </Form>

        <div className="text-center text-gray-500 text-sm mt-4">
          <p>默认管理员账号：</p>
          <p>用户名：admin</p>
          <p>密码：admin123</p>
        </div>
      </Card>
    </div>
  );
} 