import './App.css'
import Chat from './components/chat/Chat'
import Signup from './components/auth/Signup'
import Signin from './components/auth/Signin'
import { Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import { BrowserRouter } from 'react-router-dom'
import { createContext, useReducer} from 'react'
import { reducer,initialState } from './reducers/useReducer'
export const UserContext = createContext()

function Router(){

  return (
    <Routes>
      <Route path="/" element={<Chat/>}></Route>
      <Route path="/signup" element={<Signup/>}></Route>
      <Route path="/signin" element={<Signin/>}></Route>
    </Routes>
  )
}



function App() {
  const [state,dispatch] = useReducer(reducer, initialState)

  return (
    <>
      <UserContext.Provider value={{state,dispatch}}>
        <BrowserRouter>
          <Navbar />
          <Router />
        </BrowserRouter>
      </UserContext.Provider>
    </>
  )
}

export default App
