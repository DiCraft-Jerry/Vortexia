import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ConfigProvider, App as AntdApp } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import { QueryClient, QueryClientProvider } from 'react-query';
import { useAuthStore } from '@/stores/authStore';
import Layout from '@/components/layout/Layout';
import Login from '@/pages/auth/Login';
import Dashboard from '@/pages/dashboard/Dashboard';
import Projects from '@/pages/projects/Projects';
import Pipelines from '@/pages/pipelines/Pipelines';
import Builds from '@/pages/builds/Builds';
import '@/assets/styles/index.css';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
    },
  },
});

// 私有路由组件
function PrivateRoute({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuthStore();
  
  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }
  
  return <Layout>{children}</Layout>;
}

// 公共路由组件
function PublicRoute({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuthStore();
  
  if (isAuthenticated) {
    return <Navigate to="/" replace />;
  }
  
  return <>{children}</>;
}

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider locale={zhCN}>
        <AntdApp>
          <Router>
            <Routes>
              {/* 公共路由 */}
              <Route
                path="/login"
                element={
                  <PublicRoute>
                    <Login />
                  </PublicRoute>
                }
              />
              
              {/* 私有路由 */}
              <Route
                path="/"
                element={
                  <PrivateRoute>
                    <Dashboard />
                  </PrivateRoute>
                }
              />
              <Route
                path="/projects"
                element={
                  <PrivateRoute>
                    <Projects />
                  </PrivateRoute>
                }
              />
              <Route
                path="/pipelines"
                element={
                  <PrivateRoute>
                    <Pipelines />
                  </PrivateRoute>
                }
              />
              <Route
                path="/builds"
                element={
                  <PrivateRoute>
                    <Builds />
                  </PrivateRoute>
                }
              />
              
              {/* 默认重定向 */}
              <Route path="*" element={<Navigate to="/" replace />} />
            </Routes>
          </Router>
        </AntdApp>
      </ConfigProvider>
    </QueryClientProvider>
  );
}

export default App; 