import axios from 'axios'
import React, { useState } from 'react'
import logo from "../assets/logo.png"
import "../style/Login.css"

export const Login = () => {
    const [isLogin, setIsLogin] = useState(true)

    const handleLogin = (e) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget)
        const alertWrongInput = document.getElementById('alertWrongInput')

        axios.post('http://localhost:8080/login', {
            Email : formData.get("email"),
            Password : formData.get("password")
        }, {
            headers: {
                'Content-Type': 'application/json'
            },
            withCredentials: true
        })
        .then(res => {
            if(res.data.message === "success"){
                alertWrongInput.classList.add("d-none");
                // localStorage.setItem("info", "true")
                // localStorage.setItem("username", res.data.user.username)
                // localStorage.setItem("email", res.data.user.email)
                window.location.href = "/"
            }
        })
        .catch(err => {
            console.log(err)
            alertWrongInput.classList.remove("d-none");
        })
    }

    return (
        <div>
            <div className="login-container">
                <div className="row justify-content-center">
                    <div className="wrapper p-3">
                        <div className="col">
                            <div className="d-flex justify-content-center">
                                <img src={logo} className="logo"></img>
                            </div>
                        </div>
                        <div className="col-lg-12 d-flex justify-content-center">
                        <form onSubmit={handleLogin} className="p-2">
                            <div id="alertWrongInput" className="alert alert-danger d-none" role="alert">
                                Incorrect username or password.
                            </div>

                            <div className="input-group mb-4 mt-4">
                                <input type="email" className="input" name="email"/>
                                <label className="placeholder">Email address</label>    
                            </div>

                            <div className="input-group mb-4">
                                <input type="password" className="input" name="password"/>
                                <label className="placeholder">Password</label>    
                            </div>

                            <div className="d-flex justify-content-around">
                                <button type="submit" className="buttonPrimary w-50">
                                {isLogin ? "Sign in" : "Sign up"}
                                </button>
                            </div>
                        </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
