import React from 'react';
import { useNavigate } from "react-router-dom";
// import './HomePage.css';

function HomePage() {
    const navigate = useNavigate();

    const handlePvPClick = () => {
        console.log("Navigate to PvP Game");
        navigate('/pvp'); // Navigate programmatically
    };

    const handleAIPlayClick = () => {
        console.log("Navigate to Play against AI");
        navigate('/ai');
    };

    const handleRulesClick = () => {
        console.log("Show Game Rules");
        navigate('/rules');
    };

    return (
        <div className="container">
            <h1>Royal Game of Ur</h1>
            <div className="button-group">
                <button onClick={handlePvPClick}>PvP</button>
                <button onClick={handleAIPlayClick}>Play against AI</button>
                <button onClick={handleRulesClick}>Rules</button>
            </div>
        </div>
    );
}

export default HomePage;
