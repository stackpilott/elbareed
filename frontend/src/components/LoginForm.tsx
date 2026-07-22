import { useState } from 'react';
import { useAuth } from '../context/AuthContext';
import { useGoogleLogin } from '@react-oauth/google';

interface TokenResponse {
  access_token: string;
}

export const LoginForm = () => {
  const { login, loginWithGoogle } = useAuth();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      await login(email, password);
    } catch (err) {
      const authError = err as Error;
      setError(authError.message);
    } finally {
      setIsLoading(false);
    }
  };

  const googleLogin = useGoogleLogin({
    onSuccess: async (tokenResponse: Omit<TokenResponse, never>) => {
      setError('');
      setIsLoading(true);
      try {
        await loginWithGoogle(tokenResponse.access_token);
      } catch (err) {
        const authError = err as Error;
        setError(authError.message);
      } finally {
        setIsLoading(false);
      }
    },
    onError: () => setError('Google Sign In popup failed'),
    scope: 'https://mail.google.com/',
  });

  return (
    <div style={{ maxWidth: '400px', margin: '100px auto', fontFamily: 'sans-serif' }}>
      <h2>Webmail Login</h2>
      {error && <div style={{ color: 'red', marginBottom: '10px' }}>{error}</div>}
      
      <div style={{ marginBottom: '20px', display: 'flex', justifyContent: 'center' }}>
        <button 
          type="button"
          onClick={() => googleLogin()}
          disabled={isLoading}
          style={{
            padding: '10px 16px',
            fontSize: '15px',
            fontWeight: 500,
            cursor: 'pointer',
            backgroundColor: '#ffffff',
            border: '1px solid #d4d4d8',
            borderRadius: '6px',
            display: 'flex',
            alignItems: 'center',
            gap: '8px'
          }}
        >
          🔑 {isLoading ? 'Authenticating...' : 'Sign in with Google'}
        </button>
      </div>

      <div style={{ textAlign: 'center', margin: '15px 0', color: '#71717a' }}>or use App Password</div>

      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '15px' }}>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          style={{ padding: '10px', fontSize: '16px' }}
        />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="App Password"
          style={{ padding: '10px', fontSize: '16px' }}
        />
        <button type="submit" disabled={isLoading} style={{ padding: '10px', fontSize: '16px', cursor: 'pointer' }}>
          {isLoading ? 'Connecting...' : 'Login with Password'}
        </button>
      </form>
    </div>
  );
};