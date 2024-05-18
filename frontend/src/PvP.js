import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';


function PvPPage() {
    const [sessionId, setSessionId] = useState('');
    const navigate = useNavigate();

    const createSession = async () => {
        const response = await fetch('/create-session');
        const newSessionId = await response.text();
        navigate(`/play?session_id=${newSessionId}&player_id=1`);
    };

    const joinSession = () => {
        navigate(`/play?session_id=${sessionId}&player_id=2`);
    };

    return (
        <div>
            <h1>PvP Game</h1>
            <p>This is the player vs. player game page!</p>
            <button onClick={createSession}>Create Session</button>
            <div>
                <input 
                    type="text" 
                    value={sessionId} 
                    onChange={e => setSessionId(e.target.value)} 
                    placeholder="Enter session ID"
                />
                <button onClick={joinSession}>Join</button>
            </div>
        </div>
    );
}

export default PvPPage;
