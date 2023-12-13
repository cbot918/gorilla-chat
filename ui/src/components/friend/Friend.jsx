import { useState, useEffect } from 'react'
import M from 'materialize-css'
function Friend(){
  const [email, setEmail] = useState("")
  const [allUsers, setAllUsers] = useState({})
  const [offlineUsers, setOfflineUsers] = useState({})
  const [onlineUsers, setOnlineUsers] = useState({})
  const [isLoadingOffline, setIsLoadingOffline] = useState(true); // Loading state
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
    return fetch("http://localhost:8088/user/online",{
      method:"get",
      headers:{
        "Content-Type": "application/json",
        "Authorization": token
      }
    })
    .then(res => {
      if (!res.ok) {
          throw new Error('request online users failed');
      }
      return res.json();
    });
    // .then(res => res.json())
    // .then(data => {
    //   setOnlineUsers(data)
    //   setIsLoadingOnline(false)
    // })
    // .catch(err=>{
    //   console.log(err)
    //   setIsLoadingOnline(false)
    // })
  }

  function getAllUsers(token){
    return fetch("http://localhost:8088/user/all",{
      method:"get",
      headers:{
        "Content-Type": "application/json",
        "Authorization": token
      }
    })
    .then(res => {
      if (!res.ok) {
          throw new Error('request all users failed');
      }
      return res.json();
    });
    // .then(res => res.json())
    // .then(data => {
    //   setAllUsers(data)
    //   setIsLoadingAll(false)
    // })
    // .catch(err=>{
    //   console.log(err)
    //   setIsLoadingAll(false)
    // })
  }

  // function getOfflineUsers(token){
  //   fetch("http://localhost:8088/user/offline",{
  //     method:"get",
  //     headers:{
  //       "Content-Type": "application/json",
  //       "Authorization": token
  //     }
  //   })
  //   .then(res => res.json())
  //   .then(data => {
  //     setOfflineUsers(data)
  //     setIsLoadingOffline(false)
  //   })
  //   .catch(err=>{
  //     console.log(err)
  //     setIsLoadingAll(false)
  //   })
  // }

  useEffect(() => {
    const token = localStorage.getItem("token");
  
    Promise.all([getOnlineUsers(token), getAllUsers(token)])
      .then(([onlineUsersData, allUsersData]) => {
        setOnlineUsers(onlineUsersData);
        setIsLoadingOnline(false)

        setAllUsers(allUsersData);
        setIsLoadingAll(false)

        console.log(onlineUsersData)
        console.log(allUsersData)

        let offlines = allUsersData.names.filter(x => !onlineUsersData.users.includes(x));
        console.log(offlines)
        setOfflineUsers(offlines)
        setIsLoadingOffline(false)
        
        // const offlineUsers = allUsersData.names.filter(name => 
        //   !onlineUsersData.some(onlineUser.users => onlineUser.id === user.id)
        // );

        // console.log("Offline Users:", offlineUsers);
      })
      .catch(error => {
        // Handle errors
        console.error("Error fetching data:", error);
      });
  }, []);

  return(
    <div>


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
            value="刷新offline" 
            onClick={()=>{
              getOfflineUsers(localStorage.getItem("token"))
            }}
          />
          <div>
            <span>({ offlineUsers.length }) offline users: [ </span>{ !isLoadingOffline && offlineUsers ?<RenderOfflineUsers offlineUsers={offlineUsers}/>:<span>is loading...</span>}<span>]</span>
          </div>
          {/* <input 
            type="button" 
            value="刷新all" 
            onClick={()=>{
              getAllUsers(localStorage.getItem("token"))
            }}
          />
          <div>
            <span>({ allUsers.count -1 }) all users: [ </span>{ !isLoadingAll && allUsers ?<RenderAllUsers allUsers={allUsers}/>:<span>is loading...</span>}<span>]</span>
          </div> */}
        </div>
      {/* 好友申請的 code, 暫時留一下 */}
      {/* <input 
        id="email" 
        type="text"
        value={email}
        onChange={(e)=>{
          setEmail(e.target.value)
        }}
      /> */}

      {/* <input 
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
      /> */}


    </div>
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

function RenderOfflineUsers({offlineUsers}){
  return (
    <>
      {offlineUsers
        .map((name, index)=>(
        <span key={index}>{name} </span>
      ))}
    </>
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

export default Friend