import {ReactNode} from "react";
import {Link, Outlet} from "react-router-dom";

export default function GuestLayout() {
    return (
        <>
            <Outlet />
        </>
    );
}