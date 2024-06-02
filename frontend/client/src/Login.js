import React from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';

function Login() {
  return (
    <div className= 'd=flez justify-content-center align-items-center bg-primary'>
        <div className= 'p-3 bg-white w-25'>
           <form action="">
                <div className='mb-3'>
                    <label htmlFor="user">User</label>
                    <input type="user" placeholder= 'Enter User' className= 'form-control'/>
                </div>
                <div className='mb-3'>
                    <label htmlFor="password">Password</label>
                    <input type="password" placeholder= 'Enter Password'  className= 'form-control'/>
                </div>
                <button className='btn btn-success'>Login</button>
           </form>
        </div>
    </div>
  )
}

export default Login