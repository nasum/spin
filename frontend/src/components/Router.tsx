import React, { useContext, useEffect } from 'react'
import { RoconRoot, useRoutes } from 'rocon/react'
import axios from 'axios'
import { toplevelRoutes } from '../routes'
import { AppHeader } from './AppHeader'
import { AuthContext, SET_NAME } from '../context/auth'

const Routes: React.FC = () => {
  return useRoutes(toplevelRoutes)
}

export const Router: React.FC = () => {
  const { dispatch } = useContext(AuthContext)

  useEffect(() => {
    axios.get('/user').then((response) => {
      dispatch({ type: SET_NAME, payload: response.data.name })
    })
  }, [])

  return (
    <RoconRoot>
      <AppHeader />
      <Routes />
    </RoconRoot>
  )
}
