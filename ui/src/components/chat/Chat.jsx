import {useState, useEffect, useContext, useRef} from 'react'
import {UserContext} from '../../App'
function Chat(){

  const ws = useRef(null);
  // const [authBeforeWs,setAuthBeforeWs] = useState("")

  function authBeforeWs( user, token ){
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
    ws.current = new WebSocket(url)
    ws.current.onopen = () => {
      console.log("socket open")
      // ws.current.send("hihi")
    }
    ws.current.onmessage = (e) => {
      console.log(e.data)
    }
    ws.current.onerror = (e) =>{
      console.log("socket error")
      console.log(e)
    }
    ws.current.onclose = ()=>{
      console.log("socket close")
    }
  }

  const [name, setName] = useState("")

  const {state,dispatch} = useContext(UserContext)
  useEffect(()=>{
    if(state){
      const user = localStorage.getItem("user")
      const token = localStorage.getItem("token")
      setName(state.name)
      authBeforeWs( user, token )
    }
  },[state])

  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState('');

  function sendMessage(event){
    event.preventDefault();

    if (inputMessage.trim() !== '') {
        // setMessages([...messages, inputMessage]);
        
        ws.current.send(inputMessage)
        setInputMessage('')
        
    }
  }




  return (
      <div>
          <div> {name} </div>
          <div>
              {messages.map((msg, index) => (
                  <p key={index}>{msg}</p>
              ))}
          </div>
          <form onSubmit={sendMessage}>
              <input
                  type="text"
                  value={inputMessage}
                  onChange={(e) => setInputMessage(e.target.value)}
              />
              <button type="submit">Send</button>
          </form>
      </div>
  );
}

export default Chat