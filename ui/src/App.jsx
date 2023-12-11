import './App.css'
import Home from './components/home/Home'
import Signup from './components/auth/Signup'
import Signin from './components/auth/Signin'
import Friend from './components/friend/Friend'
import { Routes, Route  } from 'react-router-dom'
import Navbar from './components/Navbar'
import { BrowserRouter } from 'react-router-dom'
import { createContext, useReducer, useEffect, useState} from 'react'
import { reducer,initialState } from './reducers/useReducer'

export const UserContext = createContext()

function Router(){

  // const ws = useRef(null);
  // const [authBeforeWs,setAuthBeforeWs] = useState("")
  

  return (
    <Routes>
      <Route path="/" element={<Home/>}></Route>
      <Route path="/signup" element={<Signup/>}></Route>
      <Route path="/signin" element={<Signin/>}></Route>
      <Route path="/friend" element={<Friend/>}></Route>
    </Routes>
  )
}



function App() {
  const [state,dispatch] = useReducer(reducer, initialState)

  const [ws, setWs] = useState(null);

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
      // ws.current.send("hihi")
    }
    ws.onmessage = (e) => {
      console.log(e.data)
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
      // setName(state.name)
      authAndConnectWS( user, token )
    }
  },[])
  

  return (
    <>
      <UserContext.Provider value={{state, dispatch, ws}}>
        <BrowserRouter>
          <Navbar />
          <Router />
        </BrowserRouter>
      </UserContext.Provider>
    </>
  )
}

export default App