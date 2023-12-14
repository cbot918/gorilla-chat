import './App.css'
import Home from './components/home/Home'
import Signup from './components/auth/Signup'
import Signin from './components/auth/Signin'
import Friend from './components/users/Users'
import { Routes, Route  } from 'react-router-dom'
import Navbar from './components/Navbar'
import { BrowserRouter, useNavigate } from 'react-router-dom'
import { createContext, useReducer, useEffect, useState, useContext} from 'react'
import { reducer,initialState} from './reducers/useReducer'

export const UserContext = createContext()

function Router(){
  const navigate = useNavigate()
  useEffect(()=>{
    const user = JSON.parse(localStorage.getItem('user'))
    if(user){
      // dispatch({type:"USER", payload:user})
    } else {
      navigate("/signin")
    }
  },[])

  return (
    <Routes>
      <Route path="/" element={<Home />}></Route>
      <Route path="/signup" element={<Signup/>}></Route>
      <Route path="/signin" element={<Signin/>}></Route>
    </Routes>
  )
}



function App() {
  const [state,dispatch] = useReducer(reducer, initialState)
  const [ws, setWs] = useState(null);
  const [unreadNotifies, setUnreadNotifies] = useState(0)
  const [unreadMessages, setUnreadMessages] = useState(0)
  const [receivedMessage, setReceivedMessage] = useState("")

  function authAndConnectWS( user, token ){
    fetch("http://localhost:8088/auth/authbeforews",{
      method:"post",
      headers:{
        "Content-Type":"application/json",
        "Authorization": token
      },
      body:JSON.stringify({
        "id":user.id,
        "email":user.email
    })
    }).then(res => res.json())
    .then(data => {
      console.log(data)
      if( data.auth === "ok"){
        setupWS(user)
      } else {
        console.log("authBeforeWS failed")
      }
    }).catch(err =>{
      console.log(err)
    })
  }

  function setupWS(u){
    const user = JSON.parse(u)
    const url = `ws://localhost:8088/ws?id=${user.id}&email=${user.email}&name=${user.name}`
    const ws = new WebSocket(url)
    ws.onopen = () => {
      console.log("socket open")
    }
    ws.onmessage = (e) => {
      setReceivedMessage(e.data)
    }
    ws.onerror = (e) =>{
      console.log("socket error")
      console.log(e)
    }
    ws.onclose = ()=>{
      console.log("socket close")
    }

    setWs(ws)

  }

  useEffect(()=>{
    const user = localStorage.getItem("user")
    const token = localStorage.getItem("token")
    if(user && token){
      authAndConnectWS( user, token )
    }
  },[])

  return (
    <>
      <UserContext.Provider value={{state, dispatch, ws, receivedMessage}}>
        <BrowserRouter>
          <Navbar unreadMessages={unreadMessages}/>
          <Router />
        </BrowserRouter>
      </UserContext.Provider>
    </>
  )
}

export default App
