import { useState, useEffect, useCallback } from 'react';
import { useAuth } from '../context/AuthContext';
import { Sidebar } from './Sidebar';
import { Toolbar } from './Toolbar';
import { MessageList } from './MessageList';
import '../index.css';

export const AppLayout = () => {
  const { logout } = useAuth();
  const [sidebarWidth, setSidebarWidth] = useState(240);
  const [isDragging, setIsDragging] = useState(false);

  const handleMouseDown = useCallback(() => {
    setIsDragging(true);
  }, []);

  const handleMouseUp = useCallback(() => {
    setIsDragging(false);
  }, []);

  const handleMouseMove = useCallback((e: MouseEvent) => {
    if (!isDragging) return;
    const newWidth = Math.min(Math.max(e.clientX, 150), 500);
    setSidebarWidth(newWidth);
  }, [isDragging]);

  useEffect(() => {
    if (isDragging) {
      document.addEventListener('mousemove', handleMouseMove);
      document.addEventListener('mouseup', handleMouseUp);
    } else {
      document.removeEventListener('mousemove', handleMouseMove);
      document.removeEventListener('mouseup', handleMouseUp);
    }
    return () => {
      document.removeEventListener('mousemove', handleMouseMove);
      document.removeEventListener('mouseup', handleMouseUp);
    };
  }, [isDragging, handleMouseMove, handleMouseUp]);

  return (
    <div style={{ display: 'flex', height: '100vh', width: '100vw', backgroundColor: 'var(--bg-main)', userSelect: isDragging ? 'none' : 'auto' }}>
      <div style={{ width: `${sidebarWidth}px`, backgroundColor: 'var(--bg-sidebar)', display: 'flex', flexDirection: 'column', flexShrink: 0 }}>
        <div style={{ padding: '16px', fontWeight: 'bold', fontSize: '18px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          Webmail Pro
          <button onClick={logout} style={{ background: 'none', border: 'none', color: 'var(--text-muted)', cursor: 'pointer' }} title="Logout">
            ⏏️logout
          </button>
        </div>
        
        <div style={{ flex: 1, overflowY: 'auto' }}>
          <Sidebar />
        </div>
      </div>

      <div 
        className="sidebar-resizer" 
        title="Drag to resize"
        onMouseDown={handleMouseDown}
      ></div>

      <div style={{ flex: 1, display: 'flex', flexDirection: 'column', minWidth: 0 }}>
        <Toolbar />
        
        <main style={{ flex: 1, padding: '24px', overflowY: 'auto' }}>
          <MessageList />
        </main>
      </div>
    </div>
  );
};