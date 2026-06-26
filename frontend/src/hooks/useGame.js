import { useState } from 'react'
import { startGame, sendAction } from '../api/game'

export function useGame() {
    const [game, setGame] = useState(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)
    const [amount, setAmount] = useState(100)
    const [initialBalance, setInitialBalance] = useState(1000)

    const initGame = async () => {
        try {
            setLoading(true)
            setError(null)

            const res = await startGame(initialBalance)
            setGame(res)
        } catch (err) {
            setError('Failed to start game')
        } finally {
            setLoading(false)
        }
    }

    const action = async (type) => {
        if (!game) return

        try {
            setLoading(true)
            setError(null)

            const res = await sendAction(game.game_id, type, amount)
            setGame(res)
        } catch (err) {
            setError('Action failed')
        } finally {
            setLoading(false)
        }
    }

    const reset = () => {
        setGame(null)
        setError(null)
        setAmount(100)
    }

    return {
        game,
        loading,
        error,
        amount,
        initialBalance,
        setAmount,
        initGame,
        action,
        reset,
        setInitialBalance,
    }
}
