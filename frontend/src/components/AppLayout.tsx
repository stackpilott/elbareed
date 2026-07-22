import { useAuth } from '../context/AuthContext';

export const AppLayout = () => {
  const { logout } = useAuth();

  return (
    <div style={{ display: 'flex', height: '100vh', width: '100vw', backgroundColor: '#f4f4f5', margin: 0, padding: 0 }}>
      <nav style={{ width: '80px', backgroundColor: '#18181b', display: 'flex', flexDirection: 'column', alignItems: 'center', padding: '20px 0' }}>
        <div style={{ width: '40px', height: '40px', backgroundColor: '#3b82f6', borderRadius: '8px', marginBottom: 'auto' }}></div>
        <button 
          onClick={logout}
          style={{ background: 'transparent', color: '#a1a1aa', border: 'none', cursor: 'pointer', marginBottom: '20px' }}
        >
          Exit
        </button>
      </nav>
      <main style={{ flex: 1, padding: '40px', overflowY: 'auto' }}>
        <h1 style={{ color: '#27272a', margin: '0 0 20px 0', fontFamily: 'sans-serif' }}>
          Your Inbox Reimagined
        </h1>
        <div style={{ backgroundColor: 'white', height: '80%', borderRadius: '12px', border: '1px dashed #d4d4d8', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          <p style={{ color: '#a1a1aa', fontFamily: 'sans-serif' }}>Waiting for UI design selection...</p>
        </div>
      </main>
    </div>
  );
};