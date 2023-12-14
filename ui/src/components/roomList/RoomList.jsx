import { useEffect,useState, useContext } from 'react'
import './roomList.css'
import { UserContext } from '../../App'
function RoomList(){

  const [rooms, setRooms] = useState([{}])
  const [chattos, setChattos] = useState([{}])
  const {state, dispatch} = useContext(UserContext)
  const [activeRoom, setActiveRoom] = useState(1);
  const [roomID, setRoomID] = useState(0)
  const [roomObj, setRoomObj] = useState([{}])

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

  function wrapRoomObj(rooms, chattos){
    console.log(rooms)
    console.log(chattos)

    const editedRooms = rooms.map(r=>({...r, "type":"room"}))
    const editedChattos = chattos.map(u=>({...u, "type":"user"}))

  }

  function getRoomList(){
    const user_id = parseInt(JSON.parse(localStorage.getItem('user')).id)
    Promise.all([getDefaultRooms(), getChattoUsers(user_id)])
      .then(([defaultRooms, chattoUsers])=>{
        wrapRoomObj(defaultRooms,chattoUsers)
        setRooms(defaultRooms)
        setChattos(chattoUsers)
      })
      .catch(err => {
        console.log(err)
      })
  }

  function getDefaultRooms(){
    return fetch("http://localhost:8088/room/default",{
      method: "get",
      headers: {
        "Content-Type":"application/json",
        "Authorization":localStorage.getItem("token")
      },
    }).then(res=>{
      if (!res.ok) { throw new Error('request default room failed')}
      return res.json()
    })
  }

  function getChattoUsers(userID){
    return fetch("http://localhost:8088/room/chatto",{
      method: "post",
      headers: {
        "Content-Type":"application/json",
        "Authorization":localStorage.getItem("token")
      },
      body: JSON.stringify({"user_id":userID})
    }).then(res=>{
      if (!res.ok) { throw new Error('request chatto users failed')}
      return res.json()
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
    getRoomList()

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