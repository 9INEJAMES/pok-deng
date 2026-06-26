import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import GamePage from './pages/GamePage.jsx'

createRoot(document.getElementById('root')).render(
    <StrictMode>
        <GamePage />
    </StrictMode>,
)
