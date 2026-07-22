import { useState } from 'react';

export const Sidebar = () => {
  const [isMoreOpen, setIsMoreOpen] = useState(false);

  const navItemStyle = {
    padding: '8px 12px',
    margin: '4px 0',
    borderRadius: '6px',
    cursor: 'pointer',
    display: 'flex',
    alignItems: 'center',
    gap: '10px',
    fontWeight: 500,
  };

  return (
    <nav style={{ padding: '16px', display: 'flex', flexDirection: 'column', gap: '20px' }}>
      <div>
        <div style={{ fontSize: '12px', fontWeight: 600, color: 'var(--text-muted)', textTransform: 'uppercase', marginBottom: '8px', letterSpacing: '0.5px' }}>
          Core
        </div>
        <div style={{ ...navItemStyle, backgroundColor: '#dbeafe', color: 'var(--primary-accent)' }}>
          <span>📥</span> Inbox <span style={{ marginLeft: 'auto', fontWeight: 'bold' }}>12</span>
        </div>
        <div style={navItemStyle}><span>⭐</span> Starred</div>
        <div style={navItemStyle}><span>📤</span> Sent</div>
        <div style={navItemStyle}><span>📝</span> Drafts</div>
      </div>

      <div style={{ height: '1px', backgroundColor: 'var(--border-light)' }}></div>

      <div>
        <button 
          onClick={() => setIsMoreOpen(!isMoreOpen)}
          style={{ ...navItemStyle, background: 'transparent', border: 'none', width: '100%', color: 'var(--text-muted)' }}
        >
          <span>{isMoreOpen ? '🔽' : '▶️'}</span> {isMoreOpen ? 'Less' : 'More Labels'}
        </button>
        
        {isMoreOpen && (
          <div style={{ paddingLeft: '24px', marginTop: '8px', display: 'flex', flexDirection: 'column', gap: '4px' }}>
            <div style={{...navItemStyle, fontSize: '14px'}}>Receipts</div>
            <div style={{...navItemStyle, fontSize: '14px'}}>Newsletters</div>
            <div style={{...navItemStyle, fontSize: '14px'}}>Projects</div>
          </div>
        )}
      </div>
    </nav>
  );
};