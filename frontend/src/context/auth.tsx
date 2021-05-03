import React, { useReducer } from 'react'

type Auth = {
  name: string
}

const auth: Auth = {
  name: '',
}

export const SET_NAME = Symbol('SET_NAME')

type ActionType = { type: typeof SET_NAME; payload: string }

export const AuthContext = React.createContext<{
  state: Auth
  dispatch: React.Dispatch<ActionType>
}>({
  state: auth,
  dispatch: () => undefined,
})

const reducer = (state: Auth, action: ActionType) => {
  switch (action.type) {
    case SET_NAME:
      return {
        ...state,
        name: action.payload,
      }
    default:
      return state
  }
}

type Props = {
  children?: React.ReactNode
}

export const AuthContextProvider: React.FC = ({ children }: Props) => {
  const [state, dispatch] = useReducer(reducer, auth)

  return (
    <AuthContext.Provider value={{ state, dispatch }}>
      {children}
    </AuthContext.Provider>
  )
}
