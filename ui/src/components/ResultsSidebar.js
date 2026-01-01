import React, { useEffect, useState } from "react";

function ResultsSidebar({ competition }) {
  const [results, setResults] = useState([]);
  const [page, setPage] = useState(0);
  const [totalCount, setTotalCount] = useState(0);
  const [loading, setLoading] = useState(true);
  const pageSize = 10;

  // Reset page when competition changes
  useEffect(() => {
    setPage(0);
  }, [competition]);

  useEffect(() => {
    async function fetchResults() {
      setLoading(true);
      try {
        const res = await fetch(
          `http://localhost:5000/v1/scores/finished/${competition}?page=${page}&limit=${pageSize}`
        );
        const data = await res.json();
        setResults(data.matches || []);
        setTotalCount(data.count || 0);
      } catch (err) {
        console.error("Error fetching results:", err);
      } finally {
        setLoading(false);
      }
    }
    fetchResults();
  }, [competition, page]);

  return (
    <aside className="results-sidebar">
      <h2>Results</h2>
      {loading ? (
        <p>Loading...</p>
      ) : totalCount === 0 ? (
        <p>No finished matches.</p>
      ) : (
        <>
          <ul>
            {results.map((m, idx) => (
              <li key={idx}>
                {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                <br />
                <small>{new Date(m.match_date).toLocaleString()}</small>
              </li>
            ))}
          </ul>
          <div className="pagination">
            <button disabled={page === 0} onClick={() => setPage(page - 1)}>
              Previous
            </button>
            <span>
              Page {page + 1} of {Math.ceil(totalCount / pageSize)}
            </span>
            <button
              disabled={page + 1 >= Math.ceil(totalCount / pageSize)}
              onClick={() => setPage(page + 1)}
            >
              Next
            </button>
          </div>
        </>
      )}
    </aside>
  );
}

export default ResultsSidebar;