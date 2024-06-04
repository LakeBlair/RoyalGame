import React, { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';
import './App.css';


function PlayPage() {
    const [searchParams] = useSearchParams();
    const sessionId = searchParams.get('session_id');
    const playerId = searchParams.get('player_id');
    const [currentPid, setCurrentPid] = useState(0);
    const [board, setBoard] = useState('');
    const [turn, setTurn] = useState('');
    const [dice, setDice] = useState('');
    const [progress1, setProgress1] = useState('');
    const [progress2, setProgress2] = useState('');
    const [winner, setWinner] = useState('');
    const [move, setMove] = useState('')
    const [moveCount, setMoveCount] = useState(0)
    const [ready, setReady] = useState(false)
    const ws = new WebSocket(`ws://localhost:8080/play?session_id=${sessionId}&player_id=${playerId}`);

    useEffect(() => {
        if (sessionId) {
            ws.onopen = () => {
                console.log('WebSocket connection established');
                ws.send(JSON.stringify({ type: 'Start' }));
            };

            ws.onmessage = (event) => {
                const receivedMessage = JSON.parse(event.data);
                console.log('Received message:', receivedMessage);
                if (receivedMessage.type == "Start_ACK") {
                    setReady(!ready);
                }
                if (receivedMessage.type == "Grid") {
                    setBoard(receivedMessage.content);
                }
                if (receivedMessage.type == "TurnStart") {
                    setTurn(receivedMessage.content);
                }
                if (receivedMessage.type == "Dice") {
                    setDice(receivedMessage.content);
                }
                if (receivedMessage.type == "Progress") {
                    if (receivedMessage.receiver == 1) {
                        setProgress1(receivedMessage.content);
                    }
                    if (receivedMessage.receiver == 2) {
                        setProgress2(receivedMessage.content);
                    }
                }
                if (receivedMessage.type == "Winner") {
                    setWinner(receivedMessage.content);
                }
                if (receivedMessage.type == "Move") {
                    setMove(receivedMessage.content);
                    setMoveCount(receivedMessage.move);
                    setCurrentPid(receivedMessage.player);
                }

            };

            ws.onerror = (error) => {
                console.log('WebSocket error:', error);
            };

            ws.onclose = () => {
                console.log('WebSocket connection closed');
            };

            return () => {
                ws.close();
            };
        }
    }, [sessionId, playerId]);

    const handlePlayerMove = (move) => {
        console.log(ws);
        if (ws) {
            console.log("WebSocket state:", ws.readyState);
            if (ws.readyState === WebSocket.OPEN) {
                console.log("Sending move:", move);
                ws.send(JSON.stringify({ type: 'Move', move: move }));
            } else {
                console.log("WebSocket is not open. Current state:", ws.readyState);
            }
        } else {
            console.log("WebSocket is not initialized");
        }
    };

    return (
        <div>
            <h1>Playing the Royal Game of Ur</h1>
            <h2><pre>{winner}</pre></h2>
            <pre>{turn}</pre>
            <pre>{dice}</pre>
            {ready && <p>Player 1 Progress:</p>}
            <pre>{progress1}</pre>
            {ready && <p>Player 2 Progress:</p>}
            <pre>{progress2}</pre>
            {!ready && <p>Session ID: {sessionId}</p>}
            <pre>{board}</pre>
            <pre>{move}</pre>
            <div>
                {currentPid == playerId && Array.from({ length: moveCount }, (_, i) => (
                    <button key={i} className="game-button" onClick={() => handlePlayerMove(i)}>
                        Move {i}
                    </button>
                ))}
            </div>
        </div>
    );
}

export default PlayPage;
