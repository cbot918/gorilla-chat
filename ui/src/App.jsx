import './App.css'

function App() {

  function setupWS(){
    const ws = new WebSocket("ws://localhost:8088/ws")
    ws.onopen = () => {
      console.log("socket open")
      ws.send("hihi")
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
  }

  return (
    <>
      HOME
     {setupWS()}
    </>
  )
}

export default App
