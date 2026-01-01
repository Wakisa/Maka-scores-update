import React, { useState } from "react";
import Header from "./components/Header";
import CompetitionSelector from "./components/CompetitionSelector";
import ScoresPanel from "./components/ScoresPanel";
import ResultsSidebar from "./components/ResultsSidebar";
import "./styles.css";

function App() {
    const [competition, setCompetition] = useState("PL");

    return (
        <div className="app">
            <Header/>
            <main className="main">
                <div className="left-panel">
                    <CompetitionSelector
                    selected={competition}
                    onChange={setCompetition}
                    />
                    <ScoresPanel competition={competition}/>
                </div>
                <ResultsSidebar competition={competition}/>
            </main>
            <footer className="footer">copyright@2025</footer>
        </div>
    );
}

export default App