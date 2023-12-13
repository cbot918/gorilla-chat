import './onlineOfflineUsers.css'

function OfflineUsers({offlineUsers}){
  return (
    <>
      {offlineUsers
        .map((name, index)=>(
        <span className="user-cursor" key={index}>{name} </span>
      ))}
    </>
  )
}

export default OfflineUsers