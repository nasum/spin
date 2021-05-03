import React, { useContext } from 'react'
import { Link } from 'rocon/react'
import { toplevelRoutes } from '../routes'
import { AuthContext } from '../context/auth'

export const AppHeader: React.FC = () => {
  const { state } = useContext(AuthContext)

  return (
    <nav className="bg-gray-800">
      <div className="max-w-7xl mx-auto">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <h1 className="text-white">Spin</h1>
            </div>
          </div>
          <div className="ml-3 relative">
            <div className="text-white">
              {state.name ? (
                <span>{state.name}</span>
              ) : (
                <Link route={toplevelRoutes._.signup}>signup</Link>
              )}
            </div>
          </div>
        </div>
      </div>
    </nav>
  )
}
