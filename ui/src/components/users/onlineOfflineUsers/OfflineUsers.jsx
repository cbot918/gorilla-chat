import './onlineOfflineUsers.css'
import { UserContext } from '../../../App'
import { useState, useEffect, useContext } from 'react'
function OfflineUsers({offlineUsers}){
  const [ chattoUser, setChattoUser ] = useState({})
  const [ chattoPostCheck, setchattoPostCheck] = useState(false)
  const { state, dispatch } = useContext(UserContext)
  function postAddChatto(userID, chattoID){
    fetch("http://localhost:8088/room/chatto/add", {
      method:"post",
      headers:{
        "Content-Type":"application/json",
        "Authorization":localStorage.getItem('token')
      },
      body:JSON.stringify({"user_id":userID, "chatto_id":chattoID})
    })
    .then(res=>res.json())
    .then(data=>{
      if(data.err){
        console.log("postAddChatto failed")
        return
      }
      setchattoPostCheck(true)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  useEffect(()=>{
    dispatch({type:"ADD_CHATTO", payload:{"chatto_id":chattoUser.user_id, "chatto_name":chattoUser.name}})
  },[chattoPostCheck])

  return (
    <>
      {offlineUsers
        .filter((user)=>user.name != JSON.parse(localStorage.getItem('user')).name)
        .map((user, index)=>(
        <span 
          className="user-cursor" 
          key={user.user_id}
          data-userid={user.user_id}
          onClick={(e) => {
            const userID = parseInt(JSON.parse(localStorage.getItem('user')).id)
            const chattoID = parseInt(e.target.getAttribute('data-userid'))
            postAddChatto(userID, chattoID)
            setChattoUser({"user_id":chattoID, "name":e.target.textContent})
          }}
        >{user.name}{' '}
        </span>
      ))}
    </>
  )
}

export default OfflineUsers