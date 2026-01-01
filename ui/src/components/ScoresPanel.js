import React, { useEffect, useState } from "react";

function ScoresPanel({ competition }) {
  const [scores, setScores] = useState([]);
  const [upcoming, setUpcoming] = useState([]);
  const [loading, setLoading] = useState(true);

  // Which dataset is currently displayed: "live" or "upcoming"
  const [active, setActive] = useState("live");

  // Pagination
  const [page, setPage] = useState(0);
  const pageSize = 10;

  // Separate counts to avoid overwriting
  const [liveCount, setLiveCount] = useState(0);
  const [upcomingCount, setUpcomingCount] = useState(0);

  // Reset page and active dataset when competition changes
  useEffect(() => {
    setPage(0);
    setActive("live");
  }, [competition]);

  // Fetch data whenever competition or page changes
  useEffect(() => {
    async function fetchScores() {
      setLoading(true);
      try {
        const res = await fetch(
          `http://localhost:5000/v1/scores/live/${competition}?page=${page}&limit=${pageSize}`
        );
        const data = await res.json();
        const liveMatches = data.matches || [];
        const liveTotal = data.count || 0;

        if (liveMatches.length > 0) {
          setScores(liveMatches);
          setLiveCount(liveTotal);
          setActive("live");
          setUpcoming([]);
        } else {
          const upcomingRes = await fetch(
            `http://localhost:5000/v1/scores/upcoming/${competition}?page=${page}&limit=${pageSize}`
          );
          const upcomingData = await upcomingRes.json();
          const upcomingMatches = upcomingData.matches || [];
          const upcomingTotal = upcomingData.count || 0;

          setUpcoming(upcomingMatches);
          setUpcomingCount(upcomingTotal);
          setScores([]);
          setActive("upcoming");
        }
      } catch (err) {
        console.error("Error fetching scores:", err);
      } finally {
        setLoading(false);
      }
    }
    fetchScores();
  }, [competition, page]);

  // Derive total pages from the active dataset
  const activeCount = active === "live" ? liveCount : upcomingCount;
  const totalPages = Math.ceil(activeCount / pageSize);

  // Clamp page if it exceeds the last page of the active dataset
  useEffect(() => {
    if (totalPages > 0 && page >= totalPages) {
      setPage(totalPages - 1);
    }
  }, [totalPages, page]);

  // Current slice to render
  const currentSlice = active === "live" ? scores : upcoming;

  return (
    <div className="scores-panel">
      <h2>Live Scores</h2>

      {loading ? (
        <p>Loading...</p>
      ) : activeCount === 0 ? (
        <p>No live or upcoming matches.</p>
      ) : (
        <>
          {active === "upcoming" && <h3>Upcoming Matches</h3>}

          {currentSlice.length === 0 ? (
            <p>No matches on this page.</p>
          ) : (
            <ul>
              {currentSlice.map((m, idx) => (
                <li key={idx}>
                  {m.home_team} {m.home_score} - {m.away_score} {m.away_team}
                  {m.match_date && (
                    <> {new Date(m.match_date).toLocaleString()}</>
                  )}
                </li>
              ))}
            </ul>
          )}

          <div className="pagination">
            <button disabled={page === 0} onClick={() => setPage(page - 1)}>
              Previous
            </button>
            <span>
              Page {page + 1} of {totalPages || 1}
            </span>
            <button
              disabled={page + 1 >= totalPages}
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