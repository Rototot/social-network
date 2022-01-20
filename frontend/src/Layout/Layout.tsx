import {Outlet} from "react-router-dom";
import ResponsiveAppBar from "./ResponsiveAppBar";

export default function Layout() {
    return (
        <>
            <Header/>
            <Outlet/>

        </>
    );
}

function Header() {
    return (
        <>
            <ResponsiveAppBar/>
        </>
    );
}