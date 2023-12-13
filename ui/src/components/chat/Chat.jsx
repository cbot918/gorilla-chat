import  { useState,useContext,useRef,useEffect } from 'react';
import './chat.css'; // Make sure to create this CSS file
import {UserContext} from '../../App'

function Chat() {
    const { ws } = useContext(UserContext);

    const [newMessage, setNewMessage] = useState('');
    const [messages, setMessages] = useState([]);
    
    const handleKeyDown = (e) => {
        const user = JSON.parse(localStorage.getItem("user"))
        const msg = JSON.stringify({
            "id":       parseInt(user.id),
            "email":    user.email,
            "name":     user.name, 
            "message":  newMessage,
            "to_user":  2
        })

        if (e.key === 'Enter') {
            setMessages([...messages, { text: newMessage, mine: true }]);
            ws.send(msg)
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
    return (
        <div className="chat-container">
            <div className="messages-container" ref={messagesContainerRef}>
                {messages.map((msg, index) => (
                    <div key={index} className={`message ${msg.mine ? 'mine' : ''}`}>
                        {msg.text}
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