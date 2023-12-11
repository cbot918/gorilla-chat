import { useState, useEffect } from 'react'
function Friend(){
  const [email, setEmail] = useState("")
  const [allUsers, setAllUsers] = useState({})
  const [isLoadingOnline, setIsLoadingOnline] = useState(true); // Loading state
  const [isLoadingAll, setIsLoadingAll] = useState(true); // Loading state

  function postData(token, targetEmail){
    console.log(token)
    console.log(targetEmail)
    fetch("http://localhost:8088/friend/add", {
      method: "post",
      headers: {
        "Content-Type": "application/json",
        "Authorization": token
      },
      body: JSON.stringify({
        targetEmail
      }),
    })
    .then((res) => res.json())
    .then((data) => {
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
      setAllUsers(data)
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
          postData(localStorage.getItem("token"), email)
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
          <span>online:[ </span><span>{}</span><span>]</span>
        </div>
        <input 
          type="button" 
          value="刷新all" 
          onClick={()=>{
            getAllUsers(localStorage.getItem("token"))
          }}
        />
        <div>
          <span>allusers:[ </span>{ !isLoadingAll && allUsers ?<RenderAllUsers allUsers={allUsers}/>:<span>is loading...</span>}<span>]</span>
        </div>
      </div>
    </div>
  )
}

function RenderAllUsers({allUsers}){
  return (
    <>
    {allUsers.names.map((name, index)=>(
      <span key={index}>{name} </span>
    ))}
    </>
    
  )
}

export default Friend