import React from "react";
import ScoresList from "./components/ScoresList";

function App() {
    return (
        <div>
            <h1>Maka Live Scores</h1>
            <ScoresList competition="PL" type="live" />
            <ScoresList competition="PD" type="finished" />
        </div>
    );
}

export default App