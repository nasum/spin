import React from 'react'
import { Path } from 'rocon/react'
import { SignUp } from './components/SignUp'

function root() {
  return <p>root</p>
}

export const toplevelRoutes = Path()
  .exact({
    action: root,
  })
  .route('login', (route) => route.action(() => <p>login</p>))
  .route('signup', (route) => route.action(() => <SignUp />))
