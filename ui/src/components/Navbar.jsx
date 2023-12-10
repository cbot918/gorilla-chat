import { Link } from 'react-router-dom'

function Navbar(){
  return(
    <>
      <Link className='nav-link' to='/'>Home</Link>
      <Link className='nav-link' to='/signup'>Signup</Link>
      <Link className='nav-link' to='/signin'>Signin</Link>
    </>
  )
}

export default Navbar