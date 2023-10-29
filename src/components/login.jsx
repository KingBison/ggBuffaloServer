import React from "react";
import "../styles/login.css"

const Login = (props) => {
    const setLoggedIn = props.setLoggedIn
    const setName = props.setName
    const name = props.name


    return(
        <div className="login-main">
            <div className="login-content">
                <h1>Enter your name here</h1>
                <input 
                    type="text" 
                    onChange={(e) => {setName(e.target.value)}} 
                    value={name}
                    className="login-input"
                />
                <button
                    onClick={()=>{
                        setLoggedIn(true)
                    }}
                    className="login-button"
                >Enter</button>
            </div>
        </div>
    )

}

export default Login;