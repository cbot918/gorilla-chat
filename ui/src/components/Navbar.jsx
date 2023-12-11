import { Link } from 'react-router-dom'
import './navbar.css'

function Navbar({ unreadMessages }){
  return(
    <>
      <Link className='nav-link' to='/'>Message({unreadMessages})</Link>
      <Link className='nav-link' to='/signup'>Signup</Link>
      <Link className='nav-link' to='/signin'>Signin</Link>
      <Link className='nav-link' to='/friend'>Friend</Link>
    </>
  )
}

export default Navbar