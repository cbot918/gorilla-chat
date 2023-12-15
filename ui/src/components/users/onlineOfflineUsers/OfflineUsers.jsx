import './onlineOfflineUsers.css'

function OfflineUsers({offlineUsers}){
  return (
    <>
      {offlineUsers
        .map((user, index)=>(
        <span 
          className="user-cursor" 
          key={user.user_id}
          data-userid={user.user_id}
          onClick={(e) => {
            console.log(typeof(e.target.getAttribute('data-userid')))
            console.log(e.target.getAttribute('data-userid'))
          }}
        >{user.name} 
        </span>
      ))}
    </>
  )
}

export default OfflineUsers