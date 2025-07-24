import { Card, Row, Col, Statistic, Typography, List, Tag, Button, Space } from 'antd';
import {
  ProjectOutlined,
  BranchesOutlined,
  BuildOutlined,
  ClockCircleOutlined,
  CheckCircleOutlined,
  ExclamationCircleOutlined,
} from '@ant-design/icons';

const { Title } = Typography;

// 模拟数据
const mockStats = {
  totalProjects: 8,
  totalPipelines: 15,
  totalBuilds: 142,
  runningBuilds: 3,
};

const mockRecentBuilds = [
  {
    id: 1,
    project: 'Web Frontend',
    pipeline: 'Main Pipeline',
    status: 'success',
    branch: 'main',
    duration: 156,
    time: '2分钟前',
  },
  {
    id: 2,
    project: 'API Backend',
    pipeline: 'CI Pipeline',
    status: 'running',
    branch: 'develop',
    duration: null,
    time: '5分钟前',
  },
  {
    id: 3,
    project: 'Mobile App',
    pipeline: 'Release Pipeline',
    status: 'failed',
    branch: 'release/v1.2',
    duration: 89,
    time: '10分钟前',
  },
];

const getStatusIcon = (status: string) => {
  switch (status) {
    case 'success':
      return <CheckCircleOutlined className="text-green-500" />;
    case 'running':
      return <ClockCircleOutlined className="text-blue-500" />;
    case 'failed':
      return <ExclamationCircleOutlined className="text-red-500" />;
    default:
      return <ClockCircleOutlined className="text-gray-500" />;
  }
};

const getStatusColor = (status: string) => {
  switch (status) {
    case 'success':
      return 'green';
    case 'running':
      return 'blue';
    case 'failed':
      return 'red';
    default:
      return 'default';
  }
};

export default function Dashboard() {
  return (
    <div className="space-y-6">
      {/* 页面标题 */}
      <div className="flex items-center justify-between">
        <Title level={2} className="mb-0">
          仪表板
        </Title>
        <Button type="primary">
          创建新项目
        </Button>
      </div>

      {/* 统计卡片 */}
      <Row gutter={[16, 16]}>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic
              title="项目总数"
              value={mockStats.totalProjects}
              prefix={<ProjectOutlined />}
              valueStyle={{ color: '#1890ff' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic
              title="流水线总数"
              value={mockStats.totalPipelines}
              prefix={<BranchesOutlined />}
              valueStyle={{ color: '#52c41a' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic
              title="构建总数"
              value={mockStats.totalBuilds}
              prefix={<BuildOutlined />}
              valueStyle={{ color: '#722ed1' }}
            />
          </Card>
        </Col>
        <Col xs={24} sm={12} lg={6}>
          <Card>
            <Statistic
              title="运行中构建"
              value={mockStats.runningBuilds}
              prefix={<ClockCircleOutlined />}
              valueStyle={{ color: '#fa8c16' }}
            />
          </Card>
        </Col>
      </Row>

      {/* 最近构建 */}
      <Row gutter={[16, 16]}>
        <Col xs={24} lg={16}>
          <Card
            title="最近构建"
            extra={
              <Button type="link">
                查看全部
              </Button>
            }
          >
            <List
              dataSource={mockRecentBuilds}
              renderItem={(item) => (
                <List.Item
                  actions={[
                    <Space key="actions">
                      <span className="text-gray-500">{item.time}</span>
                      {item.duration && (
                        <span className="text-gray-500">
                          {item.duration}s
                        </span>
                      )}
                    </Space>
                  ]}
                >
                  <List.Item.Meta
                    avatar={getStatusIcon(item.status)}
                    title={
                      <Space>
                        <span>{item.project}</span>
                        <span className="text-gray-400">·</span>
                        <span className="text-gray-600">{item.pipeline}</span>
                      </Space>
                    }
                    description={
                      <Space>
                        <Tag color={getStatusColor(item.status)}>
                          {item.status}
                        </Tag>
                        <span className="text-gray-500">
                          分支: {item.branch}
                        </span>
                      </Space>
                    }
                  />
                </List.Item>
              )}
            />
          </Card>
        </Col>

        <Col xs={24} lg={8}>
          <Card title="快速操作">
            <Space direction="vertical" className="w-full">
              <Button type="primary" block icon={<ProjectOutlined />}>
                创建项目
              </Button>
              <Button block icon={<BranchesOutlined />}>
                配置流水线
              </Button>
              <Button block icon={<BuildOutlined />}>
                触发构建
              </Button>
            </Space>
          </Card>
        </Col>
      </Row>
    </div>
  );
} 