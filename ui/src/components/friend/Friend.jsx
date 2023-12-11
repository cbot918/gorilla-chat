import { useState, useEffect } from 'react'
import M from 'materialize-css'
function Friend(){
  const [email, setEmail] = useState("")
  const [allUsers, setAllUsers] = useState({})
  const [onlineUsers, setOnlineUsers] = useState({})
  const [isLoadingOnline, setIsLoadingOnline] = useState(true); // Loading state
  const [isLoadingAll, setIsLoadingAll] = useState(true); // Loading state

  function addFriendRequest(token, from, name){
    // M.toast({html: "送出邀請",classes:"#c62828 red darken-3"})
    fetch("http://localhost:8088/friend/add", {
      method: "post",
      headers: {
        "Content-Type": "application/json",
        "Authorization": token
      },
      body: JSON.stringify({
        from,
        name
      }),
    })
    .then(res=>res.json())
    .then((data) => {
      if(data.error){
        M.toast({html: data.error,classes:"#c62828 red darken-3"})
      } else {
        M.toast({html:data.msg,classes:"#43a047 green darken-1"})
      }
      console.log(data)
    })
    .catch((err) => {
      console.log(err);
    });
  }



  function getOnlineUsers(token){
    fetch("http://localhost:8088/user/online",{
      method:"get",
      headers:{
        "Content-Type": "application/json",
        "Authorization": token
      }
    })
    .then(res => res.json())
    .then(data => {
      setOnlineUsers(data)
      setIsLoadingOnline(false)
    })
    .catch(err=>{
      console.log(err)
      setIsLoadingOnline(false)
    })
  }

  function getAllUsers(token){
    fetch("http://localhost:8088/user/all",{
      method:"get",
      headers:{
        "Content-Type": "application/json",
        "Authorization": token
      }
    })
    .then(res => res.json())
    .then(data => {
      setAllUsers(data)
      setIsLoadingAll(false)
    })
    .catch(err=>{
      console.log(err)
      setIsLoadingAll(false)
    })
  }


  useEffect(()=>{
    getOnlineUsers(localStorage.getItem("token"))
    getAllUsers(localStorage.getItem("token"))
  },[])

  return(
    <div>
      <input 
        id="email" 
        type="text"
        value={email}
        onChange={(e)=>{
          setEmail(e.target.value)
        }}
      />

      <input 
        id="" 
        type="button"
        value="好友申請"
        onClick={()=>{
          const user = JSON.parse(localStorage.getItem("user"))
          addFriendRequest(
            localStorage.getItem("token"), 
            user.name,
            email
            )
          setEmail('')
        }}
      />

      <div>
      <input 
          type="button" 
          value="刷新online" 
          onClick={()=>{
            getOnlineUsers(localStorage.getItem("token"))
          }}
        />
        <div>
          <span>({ onlineUsers.count -1 }) online users: [ </span>{ !isLoadingOnline && onlineUsers ?<RenderOnlineUsers onlineUsers={onlineUsers}/>:<span>is loading...</span>}<span>]</span>
        </div>
        <input 
          type="button" 
          value="刷新all" 
          onClick={()=>{
            getAllUsers(localStorage.getItem("token"))
          }}
        />
        <div>
          <span>({ allUsers.count -1 }) all users: [ </span>{ !isLoadingAll && allUsers ?<RenderAllUsers allUsers={allUsers}/>:<span>is loading...</span>}<span>]</span>
        </div>
      </div>
    </div>
  )
}

function RenderAllUsers({allUsers}){
  return (
    <>
      {allUsers.names
        .filter((name)=> name != JSON.parse(localStorage.getItem('user')).name)
        .map((name, index)=>(
        <span key={index}>{name} </span>
      ))}
    </>
    
  )
}

function RenderOnlineUsers({onlineUsers}){
  return (
    <>
      {onlineUsers.users
        .filter((user)=>user != JSON.parse(localStorage.getItem('user')).name)
        .map((name, index)=>(
        <span key={index}>{name} </span>
      ))}
    </>
    
  )
}

export default Friend