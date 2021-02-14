import React from 'react'
import { AppHeader } from './AppHeader'
import { Router } from './Router'

export const App: React.FC = () => {
  return (
    <div className="relative bg-white">
      <AppHeader />
      <Router />
    </div>
  )
}
