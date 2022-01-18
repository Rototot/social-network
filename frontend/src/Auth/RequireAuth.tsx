import {useLocation, Navigate} from "react-router-dom";
import React from "react";


let AuthContext = React.createContext<AuthContextType>(null!);

export interface AuthContextType {
    user: {
        id: number
    };
    signin: (user: string, callback: VoidFunction) => void;
    signout: (callback: VoidFunction) => void;
}
//
// function AuthProvider({ children }: { children: React.ReactNode }) {
//     let [user, setUser] = React.useState<any>(null);
//
//     let signin = (newUser: string, callback: VoidFunction) => {
//         return fakeAuthProvider.signin(() => {
//             setUser(newUser);
//             callback();
//         });
//     };
//
//     let signout = (callback: VoidFunction) => {
//         return fakeAuthProvider.signout(() => {
//             setUser(null);
//             callback();
//         });
//     };
//
//     let value = { user, signin, signout };
//
//     return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
// }


export function RequireAuth({ children }: { children: JSX.Element }) {
    let auth = useAuth();
    let location = useLocation();

    console.log("RequireAuth")

    if (!auth || !auth.user) {
        console.log("!auth.user")

        // Redirect them to the /login page, but save the current location they were
        // trying to go to when they were redirected. This allows us to send them
        // along to that page after they login, which is a nicer user experience
        // than dropping them off on the home page.
        return <Navigate to="/signin" state={{ from: location }} replace />;
    }
    console.log(children)
    console.log("children")


    return children;
}

export function useAuth() : AuthContextType {
    return React.useContext(AuthContext);
}
