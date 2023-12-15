import { useState, useContext, useEffect } from 'react'
import { useNavigate } from "react-router-dom";
import { UserContext } from '../../App'

function Signin(){
  const {state, dispatch} = useContext(UserContext)
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const navigate = useNavigate();
  const [name, setName] = useState("")
  
  function postData(email, password){
    fetch("http://localhost:8088/auth/signin",{
      method: "post",
      headers: {
        "Content-Type":"application/json"
      },
      body: JSON.stringify({
        email,
        password
      })
    }).then(res=>res.json())
    .then(data=>{
      if(data.error){
        console.log(data.error)
      } else {
        const user = {"id":data.id, "email":data.email, "name":data.name}
        const token = data.token
        localStorage.setItem("user",JSON.stringify(user))
        localStorage.setItem("token",token)
  
        dispatch({type:"USER",payload:{user, token}})
        navigate('/')
      }

    }).catch(err=>{
      console.log(err)
    })
  }

  useEffect(()=>{
    if(state){
      console.log(state)
      setEmail(state.email)
      setPassword(state.password)
    }
  },[])

  return(
    <div>
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
        postData(email, password)
      }}/>
    </div>
  )
}

export default Signin