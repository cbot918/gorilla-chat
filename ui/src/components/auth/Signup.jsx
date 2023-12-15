import { useState, useContext, useEffect } from 'react'
import { useNavigate } from "react-router-dom";
import { UserContext } from '../../App'
function Signup(){
  
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const { state, dispatch } = useContext(UserContext)

  const navigate = useNavigate();

  function postData(name, email, password) {
  fetch("http://localhost:8088/auth/signup", {
    method: "post",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name,
      email,
      password,
    }),
  })
    .then((res) => res.json())
    .then((data) => {
      if(data.error){
        console.log(data.error)
      } else {
        dispatch({type:"SIGNIN", payload:{email,password}})
        navigate('/signin')
      }
      console.log(data);
    })
    .catch((err) => {
      console.log(err);
    });
}

  return(
    <div>
      <div>
        Name: <input 
                  id="name" 
                  type="text"
                  value={name}
                  onChange={(e)=>{
                    setName(e.target.value)
                  }}
                /> 
      </div>
      <div>
        Email: <input 
                  id="email" 
                  type="text"
                  value={email}
                  onChange={(e)=>{
                    setEmail(e.target.value)
                  }}
                />  
      </div>
      <div>
        Password: <input 
                    id="password" 
                    type="text"
                    value={password}
                    onChange={(e)=>{
                      setPassword(e.target.value)
                    }}
                  />
      </div>
      <input type="button" value="signin" onClick={()=>{
        postData(name, email, password)
        setName(''); setEmail(''); setPassword('')
      }}/>
    </div>
  )
}

export default Signup