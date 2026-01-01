import React, {useEffect, useState} from "react";

function ScoresPanel({competition}) {
    const [scores, setScores] = useState([]);
    const [upcoming, setUpcoming] = useState([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(0);
    const pageSize = 10;

    useEffect(() => {
        async function fetchScores() {
            try{
                const res = await fetch(
                  `http://localhost:5000/v1/scores/live/${competition}`
                );
                const data = await res.json();
                const liveMatches = data.matches || [];
                setScores(liveMatches);
                setPage(0);
                
                // If no live matches, fetch upcoming
                if (liveMatches.length === 0){
                    const upcomingRes = await fetch (
                      `http://localhost:5000/v1/scores/upcoming/${competition}`  
                    );
                    const upcomingData = await upcomingRes.json();
                    setUpcoming(upcomingData.matches || []);
                }
            }catch (err) {
                console.error(err);
            }finally {
                setLoading(false);
            }
            
        }
        fetchScores();
    }, [competition]);

    const paginatedScores = scores.slice(page * pageSize, (page +1) * pageSize);

    return (
        <div className="scores-panel">
            <h2>Live Scores</h2>
            {loading ? (
                <p>Loading...</p>
            ) : scores.length === 0 ? (
                upcoming.length === 0 ? (
                    <p>No live or upcoming matches.</p>
                ) : (        
                <>
                    <h3>Upcoming Matches</h3>
                <ul>
                    {upcoming.map((m, idx) => (
                        <li key={idx}>
                            {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                            {new Date(m.match_date).toLocaleString()}
                        </li>
                    ))}
                </ul>
                    </>
                )
            ) : (
                <>
                <ul>
                    {paginatedScores.map((m, idx) => (
                        <li key={idx}>
                            {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                        </li>
                    ))}
                </ul>
            
            {/*Pagination controls */}
                <div className="pagination">
                    <button disabled={page === 0} onClick={() => setPage(page - 1)}>
                        Previous
                    </button>
                    <span>
                        Page {page + 1} of {Math.ceil(scores.length / pageSize)}
                    </span>
                    <button
                    disabled={(page + 1) * pageSize >= scores.length}
                    onClick={() => setPage(page + 1)}
                    >
                        Next
                    </button>
                    </div>                
                </>
            )}
        </div>
    );
}

export default ScoresPanel;