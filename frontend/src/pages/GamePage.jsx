import { useGame } from '../hooks/useGame'

export default function GamePage() {
    const { game, loading, error, amount, initialBalance, setAmount, initGame, action, setInitialBalance } = useGame()

    if (!game) {
        return (
            <div>
                <h1>Pok Deng</h1>

                <div>
                    <label>Initial Balance (min 100)</label>
                    <input type="number" value={initialBalance} onChange={(e) => setInitialBalance(Number(e.target.value))} />
                </div>

                {initialBalance < 100 && <p style={{ color: 'red' }}>Minimum balance is 100</p>}

                <button onClick={() => initGame(initialBalance)} disabled={loading || initialBalance < 100}>
                    Start Game
                </button>

                {error && <p style={{ color: 'red' }}>{error}</p>}
            </div>
        )
    }

    return (
        <div>
            <h2>Balance: {game.balance}</h2>
            <h3>State: {game.state}</h3>
            <h3>Winner: {game.winner ?? '-'}</h3>

            {loading && <p>Loading...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            <hr />

            <h3>Player</h3>
            {game.player_hand?.map((c, i) => (
                <div key={i}>
                    {c.rank} {c.suit}
                </div>
            ))}

            <h3>Dealer</h3>
            {game.dealer_hand_visible?.map((c, i) => (
                <div key={i}>
                    {c.rank} {c.suit}
                </div>
            ))}

            <hr />

            <input type="number" value={amount} onChange={(e) => setAmount(Number(e.target.value))} />

            <div style={{ display: 'flex', gap: 8 }}>
                <button disabled={game.state !== 'WAITING_FOR_CUT'} onClick={() => action('cut')}>
                    Cut
                </button>

                <button disabled={game.state !== 'WAITING_FOR_BET'} onClick={() => action('bet')}>
                    Bet
                </button>

                <button disabled={game.state !== 'WAITING_FOR_DECISION'} onClick={() => action('draw')}>
                    Draw
                </button>

                <button disabled={game.state !== 'WAITING_FOR_DECISION'} onClick={() => action('stay')}>
                    Stay
                </button>

                <button disabled={game.state !== 'ROUND_END'} onClick={() => action('next_round')}>
                    Next Round
                </button>
            </div>
        </div>
    )
}
