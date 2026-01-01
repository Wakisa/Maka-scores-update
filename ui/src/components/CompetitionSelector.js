import React from "react";

function CompetitionSelector({selected, onChange}){
    const options = [
    { code: "PL", name: "Premier League" },
    { code: "CHA", name: "Championship" },
    { code: "PD", name: "La Liga" },
    { code: "BL1", name: "Bundesliga" },
    { code: "SA", name: "Serie A" },
    { code: "FL1", name: "Ligue 1" },
    { code: "DED", name: "Eredivisie" },
    { code: "PPL", name: "Primeira Liga" },
    { code: "CL", name: "Champions League" },
    { code: "WC", name: "World Cup" },
    { code: "EC", name: "European Championship" },
    { code: "BSA", name: "Brazil Serie A" },

    ];
    return (
        <div className="selector">
            <label htmlFor="competition">Select Competition:</label>
            <select
            id="competition"
            value={selected}
            onChange={(e) => onChange(e.target.value)}
            >
                {options.map((opt) => (
                    <option key={opt.code} value={opt.code}>
                        {opt.name}
                    </option>
                ))}
            </select>
        </div>
    );
}

export default CompetitionSelector;