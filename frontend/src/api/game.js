const BASE_URL = 'http://localhost:8080/api/v1'

export async function startGame(initial_balance) {
    const res = await fetch(`${BASE_URL}/game/start`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ initial_balance }),
    })
    return res.json()
}

export async function sendAction(gameId, action, amount) {
    const res = await fetch(`${BASE_URL}/game/${gameId}/action`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ action, amount }),
    })

    return res.json()
}
