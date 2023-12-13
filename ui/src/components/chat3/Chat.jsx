import {useState, useContext} from 'react'
import {UserContext} from '../../App'
function Chat(){

  const { ws } = useContext(UserContext);

  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState('');

  function sendMessage(event){
    event.preventDefault();

    if (inputMessage.trim() !== '') {
        ws.send(inputMessage)
        setInputMessage('')        
    }
  }

  return (
      <div>
          <div> {name} </div>
          <div>
              {messages.map((msg, index) => (
                  <p key={index}>{msg}</p>
              ))}
          </div>
          <form onSubmit={sendMessage}>
              <input
                  type="text"
                  value={inputMessage}
                  onChange={(e) => setInputMessage(e.target.value)}
              />
              <button type="submit">Send</button>
          </form>
      </div>
  );
}

export default Chat