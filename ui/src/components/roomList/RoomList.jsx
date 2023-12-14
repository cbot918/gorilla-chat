import { useEffect,useState, useContext } from 'react'
import './roomList.css'
import { UserContext } from '../../App'
function RoomList(){

  const [rooms, setRooms] = useState([{}])
  const {state, dispatch} = useContext(UserContext)
  const [activeRoom, setActiveRoom] = useState(1);
  const [roomID, setRoomID] = useState(0)

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

  function setCurrentRoom(roomID){
    const user = JSON.parse(localStorage.getItem('user'))
    setActiveRoom(roomID);
    setRoomID(roomID)
    const reqData = {"user_id":parseInt(user.id), "name":user.name, "room_id":roomID}
    enterRoomRequest(reqData)
  }

  useEffect(()=>{
    getDefaultRooms()

    // 寫死一開始到 room 大廳
    setCurrentRoom(1)
    dispatchRoomData({room_id: 1, room_name: '大廳'})
  },[])

  return(
    <div>
      {
        rooms.map((room,index)=>{
          return (
            <div 
              className={`roomlist-cursor ${activeRoom === room.room_id ? 'active' : ''}`} 
              key={index}
              onClick={()=>{
                setCurrentRoom(room.room_id)
                dispatchRoomData(room)
              }}
            >ooo {room.room_name}</div>
          )
        })
      }
    </div>
  )
}

export default RoomList