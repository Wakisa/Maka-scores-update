import React, {useEffect, useState} from "react";

function ScoresList({competition, type}) {
    const [scores, setScores] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        async function fetchScores() {
            try {
                const res = await fetch(
                    `http://localhost:5000/v1/scores/${type}/${competition}`
                );
                const data = await res.json();
                setScores(data.matches || []);
            } catch (err) {
                console.error("Error fetching scores:", err);
            } finally {
                setLoading(false);
            }
        }
        fetchScores();
    }, [competition, type]);

    if (loading) return <p>Loading...</p>;

    return (
        <div>
            <h2>{competition} {type} scores</h2>
            {scores.length === 0 ? (
                <p>No matches found.</p>
            ) : (
                <ul>
                    {scores.map((m, idx) => (
                        <li key={idx}>
                            {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                            ({m.status} on {new Date(m.match_date).toLocaleString()})
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}

export default ScoresList;