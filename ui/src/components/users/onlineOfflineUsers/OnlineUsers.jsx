import './onlineOfflineUsers.css'
import { UserContext } from '../../../App'
import { useState, useContext, useEffect } from 'react'
function OnlineUsers({onlineUsers}){

  const { state, dispatch } = useContext(UserContext)
  const [ chattoUser, setChattoUser ] = useState({})
  const [ chattoPostCheck, setchattoPostCheck] = useState(false)
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
    console.log("in useEffect")
    console.log(chattoUser)
    dispatch({type:"ADD_CHATTO", payload:{"chatto_id":chattoUser.user_id, "chatto_name":chattoUser.name}})
  },[chattoPostCheck])

  return (
    <>
      {
        // console.log(onlineUsers)
        onlineUsers
        .filter((user)=>user.name != JSON.parse(localStorage.getItem('user')).name)
        .map(user=>(
          <span
            className="user-cursor"
            key={user.user_id}
            data-userid={user.user_id}
            onClick={(e) => {
              console.log(typeof(e.target.getAttribute('data-userid')))
              console.log(e.target.getAttribute('data-userid'))
              const userID = parseInt(JSON.parse(localStorage.getItem('user')).id)
              const chattoID = parseInt(e.target.getAttribute('data-userid'))
              postAddChatto(userID, chattoID)
              console.log(chattoID)
              setChattoUser({"user_id":chattoID, "name":e.target.textContent})
            }}
          >
            {user.name}
          </span>
        ))
      // onlineUsers
      //   .filter((name)=>name != JSON.parse(localStorage.getItem('user')).name)
      //   .map((name, index)=>(
      //   <span 
      //     className="user-cursor" 
      //     key={index}
      //     onClick={(e) => {
      //       console.log(e.target.value)
      //       // const userID = parseInt(JSON.parse(localStorage.getItem('user')).id)
      //       // postAddChatto(userID, )
      //     }}
      //   >{name} </span>
      // ))
      }
    </>
  )
}

export default OnlineUsers