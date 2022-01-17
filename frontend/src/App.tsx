import React from 'react';
import {Route, Routes} from "react-router-dom";
import './App.css';
import SignIn from "./Auth/Signin";
import SignUp from "./Auth/Signup";

function App() {
    return (
        <div className="App">
            <header className="App-header">
            </header>
            <SignIn/>
            <SignUp/>
            <Routes>
                <Route path="/auth/login"/>
                <Route path="/auth/registration"/>
                <Route path="/auth/logout"/>
                <Route path="/messages"/>
            </Routes>
        </div>
    );
}

export default App;
