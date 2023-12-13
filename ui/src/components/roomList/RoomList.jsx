import { useEffect,useState, useContext } from 'react'
import './roomList.css'
import { UserContext } from '../../App'
function RoomList(){

  const [rooms, setRooms] = useState([{}])
  const {state, dispatch} = useContext(UserContext)
  const [activeRoom, setActiveRoom] = useState(1);

  useEffect(()=>{
    getDefaultRooms()
  },[])

  function dispatchRoomData(roomData){
    dispatch({type:"ROOM",payload:roomData})
  }
  function setListRoomColor(){
    
  }

  function getDefaultRooms(){
    fetch("http://localhost:8088/room/default",{
      method: "get",
      headers: {
        "Content-Type":"application/json"
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
                dispatchRoomData(room)
                setActiveRoom(room.room_id);
              }}
            >ooo {room.room_name}</div>
          )
        })
      }
    </div>
  )
}

export default RoomList