import React from 'react'
import { RoconRoot, useRoutes } from 'rocon/react'
import { toplevelRoutes } from '../routes'
import { AppHeader } from './AppHeader'

const Routes: React.FC = () => {
  return useRoutes(toplevelRoutes)
}

export const Router: React.FC = () => {
  return (
    <RoconRoot>
      <AppHeader />
      <Routes />
    </RoconRoot>
  )
}
