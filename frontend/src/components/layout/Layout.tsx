import { useState } from 'react';
import { Layout as AntdLayout, Menu, Avatar, Dropdown, Typography, Button } from 'antd';
import { useNavigate, useLocation } from 'react-router-dom';
import {
  DashboardOutlined,
  ProjectOutlined,
  BranchesOutlined,
  BuildOutlined,
  UserOutlined,
  LogoutOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
} from '@ant-design/icons';
import { useAuthStore } from '@/stores/authStore';
import type { MenuProps } from 'antd';

const { Header, Sider, Content } = AntdLayout;
const { Title } = Typography;

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const [collapsed, setCollapsed] = useState(false);
  const { user, logout } = useAuthStore();
  const navigate = useNavigate();
  const location = useLocation();

  // 菜单项
  const menuItems: MenuProps['items'] = [
    {
      key: '/',
      icon: <DashboardOutlined />,
      label: '仪表板',
    },
    {
      key: '/projects',
      icon: <ProjectOutlined />,
      label: '项目管理',
    },
    {
      key: '/pipelines',
      icon: <BranchesOutlined />,
      label: '流水线',
    },
    {
      key: '/builds',
      icon: <BuildOutlined />,
      label: '构建历史',
    },
  ];

  // 用户下拉菜单
  const userMenuItems: MenuProps['items'] = [
    {
      key: 'profile',
      icon: <UserOutlined />,
      label: '个人信息',
    },
    {
      type: 'divider',
    },
    {
      key: 'logout',
      icon: <LogoutOutlined />,
      label: '退出登录',
      danger: true,
    },
  ];

  const handleMenuClick = ({ key }: { key: string }) => {
    navigate(key);
  };

  const handleUserMenuClick = ({ key }: { key: string }) => {
    if (key === 'logout') {
      logout();
    } else if (key === 'profile') {
      // TODO: 打开个人信息模态框
    }
  };

  return (
    <AntdLayout className="min-h-screen">
      {/* 侧边栏 */}
      <Sider
        trigger={null}
        collapsible
        collapsed={collapsed}
        className="bg-white shadow-lg"
        width={240}
      >
        <div className="p-4 text-center border-b">
          <Title level={4} className="mb-0 text-blue-600">
            {collapsed ? 'CI' : 'Simple CI/CD'}
          </Title>
        </div>
        
        <Menu
          mode="inline"
          selectedKeys={[location.pathname]}
          items={menuItems}
          onClick={handleMenuClick}
          className="border-r-0"
        />
      </Sider>

      {/* 主体内容 */}
      <AntdLayout>
        {/* 顶部导航 */}
        <Header className="bg-white px-4 shadow-sm flex items-center justify-between">
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            className="text-lg"
          />

          <div className="flex items-center space-x-4">
            <span className="text-gray-600">
              欢迎回来，{user?.username}
            </span>
            
            <Dropdown
              menu={{
                items: userMenuItems,
                onClick: handleUserMenuClick,
              }}
              placement="bottomRight"
            >
              <Avatar
                icon={<UserOutlined />}
                className="cursor-pointer bg-blue-500"
              />
            </Dropdown>
          </div>
        </Header>

        {/* 内容区域 */}
        <Content className="p-6 bg-gray-50 min-h-full">
          {children}
        </Content>
      </AntdLayout>
    </AntdLayout>
  );
} 