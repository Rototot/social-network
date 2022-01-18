import {ReactNode} from "react";
import {Link, Outlet} from "react-router-dom";

export default function Layout() {
    return (
        <>
            <Sidebar/>
            <Header/>
            <Outlet />
            <Footer/>

        </>
    );
}

function MainWrapper(props: { children: ReactNode }) {
    return props.children;
}

// function MainContent() {
//     return (
//
//     );
// }


function Sidebar() {
    return (
        <nav>
            Sidebar
        </nav>
    )
}

function Header() {
    return (
        <>
            <header>
                This is header
            </header>
            <div>
                {/* A "layout route" is a good place to put markup you want to
          share across all the pages on your site, like navigation. */}
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Home</Link>
                        </li>
                        <li>
                            <Link to="/signin">SignIn</Link>
                        </li>
                        <li>
                            <Link to="/signup">SingUp</Link>
                        </li>
                        <li>
                            <Link to="/user-card">User Card</Link>
                        </li>
                    </ul>
                </nav>

                <hr/>

                {/* An <Outlet> renders whatever child route is currently active,
          so you can think about this <Outlet> as a placeholder for
          the child routes we defined above. */}
            </div>
        </>
    );
}

function Footer() {
    return (
        <footer>
            This is footer
        </footer>
    );
}