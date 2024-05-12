import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './HomePage';
import PvPPage from './PvP';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Router>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/pvp" element={<PvPPage />} />
          </Routes>
        </Router>
      </header>
    </div>
  );
}

export default App;
