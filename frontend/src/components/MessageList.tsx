interface Email {
  id: string;
  sender: string;
  subject: string;
  snippet: string;
  date: string;
  isUnread: boolean;
  isPinned: boolean;
}

export const MessageList = () => {
  const mockEmails: Email[] = [
    { id: '1', sender: 'GitHub', subject: 'Action Required: Security update', snippet: 'Dependabot found a vulnerability in your repository.', date: '10:42 AM', isUnread: true, isPinned: true },
    { id: '2', sender: 'Stripe', subject: 'Payment received', snippet: 'You just received a payment of $499.00.', date: '09:15 AM', isUnread: true, isPinned: true },
    { id: '3', sender: 'Alice Johnson', subject: 'Updated Q3 Roadmap', snippet: 'Hey, I made the changes we discussed in the meeting yesterday.', date: 'Yesterday', isUnread: true, isPinned: false },
    { id: '4', sender: 'AWS Billing', subject: 'Your AWS Invoice', snippet: 'Your invoice for the period of June is now available.', date: 'Jul 20', isUnread: false, isPinned: false },
    { id: '5', sender: 'Spotify', subject: 'New logins to your account', snippet: 'We noticed a new login from a device in Sweden.', date: 'Jul 19', isUnread: false, isPinned: false },
  ];

  const pinnedEmails = mockEmails.filter(email => email.isPinned);
  const standardEmails = mockEmails.filter(email => !email.isPinned);

  const renderEmailRow = (email: Email) => (
    <div key={email.id} className={`message-row ${email.isUnread ? 'unread' : 'read'}`} tabIndex={0}>
      <div style={{ width: '40px', display: 'flex', justifyContent: 'center' }}>
        <input type="checkbox" style={{ cursor: 'pointer' }} />
      </div>
      <div style={{ width: '30px', color: email.isPinned ? '#eab308' : 'var(--border-light)' }}>
        {email.isPinned ? '📌' : '☆'}
      </div>
      <div style={{ width: '180px', flexShrink: 0 }}>
        {email.sender}
      </div>
      <div style={{ flex: 1, display: 'flex', overflow: 'hidden', alignItems: 'baseline' }}>
        <span>{email.subject}</span>
        <span className="message-snippet">- {email.snippet}</span>
      </div>
      <div style={{ width: '80px', textAlign: 'right', fontSize: '13px', color: email.isUnread ? 'var(--primary-accent)' : 'var(--text-muted)' }}>
        {email.date}
      </div>
    </div>
  );

  return (
    <div style={{ display: 'flex', flexDirection: 'column', border: '1px solid var(--border-light)', borderRadius: '8px', overflow: 'hidden', backgroundColor: 'var(--bg-main)' }}>
      {pinnedEmails.length > 0 && (
        <div>
          <div style={{ padding: '8px 16px', fontSize: '12px', fontWeight: 600, color: 'var(--text-muted)', backgroundColor: 'var(--bg-sidebar)', textTransform: 'uppercase' }}>
            Pinned
          </div>
          {pinnedEmails.map(renderEmailRow)}
        </div>
      )}

      <div>
        <div style={{ padding: '8px 16px', fontSize: '12px', fontWeight: 600, color: 'var(--text-muted)', backgroundColor: 'var(--bg-sidebar)', textTransform: 'uppercase' }}>
          Everything Else
        </div>
        {standardEmails.map(renderEmailRow)}
      </div>
    </div>
  );
};