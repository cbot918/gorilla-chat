import RoomList from '../roomList/RoomList'
import Chat from '../chat/Chat'
import Users from '../users/Users'
import './home.css'
function Home(){

  return(
    <div className="home-container">
      <div className="chattop-container">
        <RoomList />
        <Chat />
      </div>
      <div className="users-container">
        <Users />
      </div>
    </div>
  )
}

export default Home