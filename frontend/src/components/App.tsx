import React from 'react'
import { Router } from './Router'
import { AuthContextProvider } from '../context/auth'

export const App: React.FC = () => {
  return (
    <AuthContextProvider>
      <div className="relative bg-gray-100 h-screen">
        <Router />
      </div>
    </AuthContextProvider>
  )
}
