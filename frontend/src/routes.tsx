import React from 'react'
import { Path } from 'rocon/react'

function root() {
  return <p>root</p>
}

export const toplevelRoutes = Path()
  .exact({
    action: root,
  })
  .route('login', (route) => route.action(() => <p>login</p>))
  .route('signin', (route) => route.action(() => <p>signin</p>))
