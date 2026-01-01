import React, {useEffect, useState} from "react";

function ResultsSidebar({competition}) {
    const [results, setResults] = useState([]);
    const [page, setPage] = useState(0);
    const pageSize = 10;

    useEffect(() => {
        async function fetchResults() {
            const res = await fetch(
                `http://localhost:5000/v1/scores/finished/${competition}`);
                const data = await res.json();
                const allMatches = data.matches || [];
            
                // Sort by date descending
                const sorted = allMatches.sort((a, b) => new Date(b.match_date) - new Date(a.match_date))
            setResults(sorted);            
        }
        fetchResults();
    }, [competition]);

    // Paginate
    const paginated = results.slice(page * pageSize, (page + 1) * pageSize);

    return (
        <aside className="results-sidebar">
            <h2>Results</h2>
            <ul>
                {paginated.map((m, idx) => (
                    <li key={idx}>
                        {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                        <br/>
                        <small>{new Date(m.match_date).toLocaleString()}</small>
                    </li>
                    ))}
            </ul>

            {/* Pagination controls go here */}
            <div className="pagination">
                <button disabled={page === 0} onClick={() => setPage(page - 1)}>Previous</button>
                <span>Page {page + 1}</span>
                <button
                disabled={(page + 1) * pageSize >= results.length}
                onClick={() => setPage(page + 1)}
                >
                    Next
                </button>
            </div>
        </aside>
    );
}

export default ResultsSidebar;