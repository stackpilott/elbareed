export const Toolbar = () => {
  return (
    <div style={{ 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'space-between',
      padding: '12px 24px',
      borderBottom: '1px solid var(--border-light)',
      backgroundColor: 'var(--bg-main)'
    }}>
      <div style={{ display: 'flex', gap: '8px' }}>
        <button className="action-btn" title="Archive (E)">
          📦 Archive
        </button>
        <button className="action-btn danger" title="Delete (#)">
          🗑️ Delete
        </button>
        <button className="action-btn" title="Move to Folder (V)">
          📁 Move
        </button>
        <button className="action-btn" title="More Actions">
          ⋮
        </button>
      </div>

      <div style={{ flex: 1, maxWidth: '500px', marginLeft: '24px' }}>
        <div style={{ 
          display: 'flex', 
          background: 'var(--bg-sidebar)', 
          borderRadius: '8px', 
          padding: '6px 12px',
          border: '1px solid var(--border-light)'
        }}>
          <span style={{ marginRight: '8px' }}>🔍</span>
          <input 
            type="text" 
            placeholder="Search mail or apply filters..." 
            style={{ border: 'none', background: 'transparent', width: '100%', outline: 'none' }}
          />
          <button style={{ border: 'none', background: 'transparent', cursor: 'pointer', color: 'var(--text-muted)' }}>
            ⚙️
          </button>
        </div>
      </div>
    </div>
  );
};