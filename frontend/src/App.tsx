import { AuthProvider, useAuth } from './context/AuthContext';
import { LoginForm } from './components/LoginForm';

const AppContent = () => {
  const { token, logout } = useAuth();

  if (!token) {
    return <LoginForm />;
  }

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Successfully Authenticated!</h1>
      <p>Your secure token: {token}</p>
      <button onClick={logout} style={{ padding: '10px' }}>Logout</button>
    </div>
  );
};

export const App = () => {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
};

export default App;