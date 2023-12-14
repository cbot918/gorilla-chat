import  { useState,useContext,useRef,useEffect } from 'react';
import './chat.css'; 
import {UserContext} from '../../App'

function Chat() {
    const { state, receivedMessage} = useContext(UserContext);
    const [room, setRoom] = useState({})

    const [newMessage, setNewMessage] = useState('');
    const [messages, setMessages] = useState([]);

    function sendMessage(messageBody){
        fetch("http://localhost:8088/message/room",{
            method: "post",
            headers: {
                "Content-Type":"application/json",
                "Authorization":localStorage.getItem("token")
            },
            body: JSON.stringify(messageBody)
            }).then(res=>res.json())
            .then(data=>{
            }).catch(err=>{
                console.log(err)
            })
    }

    const handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            const user = JSON.parse(localStorage.getItem("user"))
            const msg = {
                "room_id": state.room_id,
                "user_id":  parseInt(user.id),
                "email":    user.email,
                "name":     user.name, 
                "message":  newMessage,
                // "to_user":  2
            }
            sendMessage(msg)
            setNewMessage(''); 
        }
    };

    const scrollToBottom = () => {
        messagesContainerRef.current.scrollTop = messagesContainerRef.current.scrollHeight;
    };
    useEffect(() => {
        scrollToBottom();
    }, [messages]); 
    const messagesContainerRef = useRef(null);

    useEffect(() => {
        // to fix: 抓掉一個ws來的空白 message, 不知道為什麼會收到
        // 這邊暫時workaround
        if(receivedMessage===""){
            return
        }
        let user = JSON.parse(localStorage.getItem('user'))
        let m = {}
        try {
            m = JSON.parse(receivedMessage);
        } catch (error) {
            console.error("Parsing error:", error);
        }

        setMessages([...messages, { name:m.name , content: m.message, mine: m.user_id === parseInt(user.id)? true:false }]);
    },[receivedMessage])

    function fetchRoomHistoryMessages(roomID){
        fetch(`http://localhost:8088/message/room/${roomID}`,{
            method: "get",
            headers: {
              "Content-Type":"application/json",
              "Authorization": localStorage.getItem('token')
            },
          }).then(res=>res.json())
          .then(data=>{
            let user_id = parseInt(JSON.parse(localStorage.getItem('user')).id)
            const updatedData = data.map(m=> ({...m, mine:user_id === m.user_id}))
            setMessages(updatedData)
          }).catch(err=>{
            console.log(err)
          })
    }

    function fetchUserHistoryMessages(userID,chattoID){
        console.log(userID)
        console.log(chattoID)
        fetch(`http://localhost:8088/message/user_history`,{
            method: "post",
            headers: {
              "Content-Type":"application/json",
              "Authorization": localStorage.getItem('token')
            },
          }).then(res=>res.json())
          .then(data=>{
            console.log(data)
            // let user_id = parseInt(JSON.parse(localStorage.getItem('user')).id)
            // const updatedData = data.map(m=> ({...m, mine:user_id === m.user_id}))
            // setMessages(updatedData)
          }).catch(err=>{
            console.log(err)
          })
    }

    useEffect(()=>{
        const roomState = state
        if(roomState){
            if(roomState.type === "room"){
                fetchRoomHistoryMessages(state.room_id)
            }
            if(roomState.type === "user"){
                const myID = parseInt(JSON.parse(localStorage.getItem('user')).id)
                const chattoID = roomState.user_id
                fetchUserHistoryMessages(myID,chattoID)
            }
        }

    },[state])

    return (
        <div className="chat-container">
            <div className="messages-container" ref={messagesContainerRef}>
                {messages.map((msg, index) => (
                    <div key={index} className={`message ${msg.mine ? 'mine' : ''}`}>
                        {msg.mine?<p>{msg.content}</p>: <p>{msg.name}: {msg.content}</p>}
                    </div>
                ))}
            </div>
            <div className="input-container">
                <input
                    type="text"
                    value={newMessage}
                    onChange={(e) => setNewMessage(e.target.value)}
                    placeholder="Type a message..."
                    onKeyDown={handleKeyDown}
                />
            </div>
        </div>
    );
}

export default Chat;