import './onlineOfflineUsers.css'

function OnlineUsers({onlineUsers}){

  return (
    <>
      {onlineUsers.users
        .filter((user)=>user != JSON.parse(localStorage.getItem('user')).name)
        .map((name, index)=>(
        <span 
          className="user-cursor" 
          key={index}
          onClick={() => {}}
        >{name} </span>
      ))}
    </>
  )
}

export default OnlineUsers