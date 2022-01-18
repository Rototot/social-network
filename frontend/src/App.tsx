import React from 'react';
import {Route, Routes,} from "react-router-dom";
import './App.css';
import SignIn from "./Auth/Signin";
import SignUp from "./Auth/Signup";
import Layout from "./Layout/Layout";
import UserCard from "./UserCard/UserCard";
import Home from "./Home/Home";
import {NotFound} from "./Error/4xx";
import {RequireAuth, useAuth} from "./Auth/RequireAuth";

function App() {
    return (
        <>
            <Routes>
                <Route path="/" element={<Layout/>}>
                    <Route index element={<Home/>}/>

                    {/* guest */}
                    <Route path="signin" element={<SignIn/>}/>
                    <Route path="signup" element={<SignUp/>}/>

                    {/* user */}
                    <Route path="user-card" element={<RequireAuth><UserCard/></RequireAuth>}/>
                    {/*<Route path="messages" element={}/>*/}
                    <Route path="*" element={<NotFound/>}/>
                </Route>
            </Routes>
        </>
    );
}

export default App;
