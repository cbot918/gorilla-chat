import { useEffect,useState, useContext } from 'react'
import './roomList.css'
import { UserContext } from '../../App'
function RoomList(){

  const [rooms, setRooms] = useState([{}])
  const {state, dispatch} = useContext(UserContext)
  const [activeRoom, setActiveRoom] = useState(1);
  const [roomID, setRoomID] = useState(0)
  useEffect(()=>{
    getDefaultRooms()
  },[])

  function dispatchRoomData(roomData){
    dispatch({type:"ROOM",payload:roomData})
  }

  function enterRoomRequest(reqData){
    // console.log(reqData)
    fetch("http://localhost:8088/room/enter",{
      method: "post",
      headers: {
        "Content-Type":"application/json",
        "Authorization":localStorage.getItem("token")
      },
      body: JSON.stringify(reqData)
    }).then(res=>res.json())
    .then(data=>{
      // console.log(data)
    }).catch(err=>{
      console.log(err)
    })
  }

  function getDefaultRooms(){
    fetch("http://localhost:8088/room/default",{
      method: "get",
      headers: {
        "Content-Type":"application/json",
        "Authorization":localStorage.getItem("token")
      },
    }).then(res=>res.json())
    .then(data=>{
      setRooms(data)
    }).catch(err=>{
      console.log(err)
    })
  }

  return(
    <div>
      {
        rooms.map((room,index)=>{
          return (
            <div 
              className={`roomlist-cursor ${activeRoom === room.room_id ? 'active' : ''}`} 
              key={index}
              onClick={()=>{
                const user = JSON.parse(localStorage.getItem('user'))
                setActiveRoom(room.room_id);
                setRoomID(room.room_id)
                dispatchRoomData(room)

                const reqData = {"user_id":parseInt(user.id), "name":user.name, "room_id":room.room_id}
                enterRoomRequest(reqData)
              }}
            >ooo {room.room_name}</div>
          )
        })
      }
    </div>
  )
}

export default RoomList