import  { useState,useContext,useRef,useEffect } from 'react';
import './chat.css'; // Make sure to create this CSS file
import {UserContext} from '../../App'

function Chat() {
    const { ws,state, } = useContext(UserContext);
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
                // todo: render received message
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
            setMessages([...messages, { content: newMessage, mine: true }]);
            sendMessage(msg)
            // ws.send(JSON.stringify(msg))
            setNewMessage(''); 
        }
    };

    const scrollToBottom = () => {
        messagesContainerRef.current.scrollTop = messagesContainerRef.current.scrollHeight;
    };
    useEffect(() => {
        scrollToBottom();
    }, [messages]); // Scrolls to bottom every time messages change
    const messagesContainerRef = useRef(null);

    function fetchMessages(roomID){
        fetch(`http://localhost:8088/message/room/${roomID}`,{
            method: "get",
            headers: {
              "Content-Type":"application/json",
              "Authorization": localStorage.getItem('token')
            },
          }).then(res=>res.json())
          .then(data=>{
            setMessages(data)
          }).catch(err=>{
            console.log(err)
          })
    }

    useEffect(()=>{
        if(state){
            console.log(state)
            setRoom(state)
            fetchMessages(state.room_id)
        }

        // renderMessages()
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
                {/* <button onClick={handleSendMessage}>Send</button> */}
            </div>
        </div>
    );
}

export default Chat;