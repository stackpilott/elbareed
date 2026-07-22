import { AuthProvider, useAuth } from './context/AuthContext';
import { LoginForm } from './components/LoginForm';
import { AppLayout } from './components/AppLayout';

const AppContent = () => {
  const { token } = useAuth();

  if (!token) {
    return <LoginForm />;
  }

  return <AppLayout />;
};

export const App = () => {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
};

export default App;