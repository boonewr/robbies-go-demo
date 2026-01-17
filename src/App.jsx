import { useState } from 'react';

export default function App() {
  const [status, setStatus] = useState("Idle");
  const [count, setCount] = useState(0);

  const pingServer = async () => {
    try {
      setStatus("Pinging...");
      const res = await fetch('/api/ping');
      const data = await res.json();
      setStatus(data.message);
      setCount(data.count);
    } catch (error) {
      setStatus("Error connecting to server");
    }
  };

  return (
    <div style={{ 
      display: 'flex', 
      flexDirection: 'column', 
      justifyContent: 'center', 
      alignItems: 'center', 
      minHeight: '100vh', 
      width: '100vw',
      fontFamily: 'sans-serif', 
      textAlign: 'center' 
    }}>
      <h1>React + Go Demo</h1>
      <div style={{ margin: '20px' }}>
        <button 
          onClick={pingServer}
          style={{ padding: '10px 20px', fontSize: '1.2rem', cursor: 'pointer' }}
        >
          Ping Server
        </button>
      </div>
      <p><strong>Server Response:</strong> {status}</p>
      <p><strong>Global Server Clicks:</strong> {count}</p>
    </div>
  );
}