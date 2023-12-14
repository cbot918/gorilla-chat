import { useState, useEffect } from 'react'
import OnlineUsers from './onlineOfflineUsers/OnlineUsers'
import OfflineUsers from './onlineOfflineUsers/OfflineUsers'
function Friend(){

  const [offlineUsers, setOfflineUsers] = useState({})
  const [onlineUsers, setOnlineUsers] = useState({})
  const [isLoadingOffline, setIsLoadingOffline] = useState(true); 
  const [isLoadingOnline, setIsLoadingOnline] = useState(true); 

  function getOnlineUsers(token){
    return fetch("http://localhost:8088/user/online",{
      method:"get",
      headers:{
        "Content-Type": "application/json",
        "Authorization": token
      }
    })
    .then(res => {
      if (!res.ok) { throw new Error('request online users failed'); }
      return res.json();
    });
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
      if (!res.ok) { throw new Error('request all users failed'); }
      return res.json();
    });
  }

  function refreshUsers(token){
    Promise.all([getOnlineUsers(token), getAllUsers(token)])
    .then(([onlineUsersData, allUsersData]) => {
      // console.log(onlineUsersData)
      // console.log(allUsersData)
      setOnlineUsers(onlineUsersData);
      setIsLoadingOnline(false)

      let offlines = allUsersData.names.filter(x => !onlineUsersData.users.includes(x));
      setOfflineUsers(offlines)
      setIsLoadingOffline(false)
    })
    .catch(error => {
      console.error("Error fetching data:", error);
    });
  }

  useEffect(() => {
    const token = localStorage.getItem("token");
    refreshUsers(token)
  }, []);

  return(
    <div>

      <div>
        <input 
            type="button" 
            value="刷新users" 
            onClick={()=>{
              refreshUsers(localStorage.getItem("token"))
            }}
          />
          <div>
            <span>({ onlineUsers.count -1 }) online users: [ </span>{ !isLoadingOnline && onlineUsers ?<OnlineUsers onlineUsers={onlineUsers}/>:<span>is loading...</span>}<span>]</span>
          </div>


          <div>
            <span>({ offlineUsers.length }) offline users: [ </span>{ !isLoadingOffline && offlineUsers ?<OfflineUsers offlineUsers={offlineUsers}/>:<span>is loading...</span>}<span>]</span>
          </div>

        </div>
    </div>
  )
}








export default Friend

// unused code

// setAllUsers(allUsersData);
// setIsLoadingAll(false)

// function RenderAllUsers({allUsers}){
//   return (
//     <>
//       {allUsers.names
//         .filter((name)=> name != JSON.parse(localStorage.getItem('user')).name)
//         .map((name, index)=>(
//         <span key={index}>{name} </span>
//       ))}
//     </>
//   )
// }

// function addFriendRequest(token, from, name){
//   // M.toast({html: "送出邀請",classes:"#c62828 red darken-3"})
//   fetch("http://localhost:8088/friend/add", {
//     method: "post",
//     headers: {
//       "Content-Type": "application/json",
//       "Authorization": token
//     },
//     body: JSON.stringify({
//       from,
//       name
//     }),
//   })
//   .then(res=>res.json())
//   .then((data) => {
//     if(data.error){
//       M.toast({html: data.error,classes:"#c62828 red darken-3"})
//     } else {
//       M.toast({html:data.msg,classes:"#43a047 green darken-1"})
//     }
//     console.log(data)
//   })
//   .catch((err) => {
//     console.log(err);
//   });
// }


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