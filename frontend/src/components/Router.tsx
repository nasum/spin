import React from 'react'
import { Path, RoconRoot, useRoutes } from 'rocon/react'

function root() {
  return <p>root</p>
}

const toplevelRoutes = Path()
  .exact({
    action: root,
  })
  .route('login', (route) => route.action(() => <p>login</p>))
  .route('signin', (route) => route.action(() => <p>signin</p>))

const Routes: React.FC = () => {
  return useRoutes(toplevelRoutes)
}

export const Router: React.FC = () => {
  return (
    <RoconRoot>
      <Routes />
    </RoconRoot>
  )
}
