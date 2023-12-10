import { useState } from 'react'

function Signup(){

  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  return(
    <div>
      Name: <input 
                id="name" 
                type="text"
                onChange={(e)=>{
                  setName(e.target.value)
                  console.log(name)
                }}
              /> 
      Email: <input 
                id="email" 
                type="text"
                onChange={(e)=>{
                  setEmail(e.target.value)
                  console.log(email)
                }}
              />  
      Password: <input 
                  id="password" 
                  type="text"
                  onChange={(e)=>{
                    setPassword(e.target.value)
                    console.log(password)
                  }}
                />
    </div>
  )
}

export default Signup