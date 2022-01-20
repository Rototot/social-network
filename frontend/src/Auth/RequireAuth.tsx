import {Navigate, useLocation} from "react-router-dom";
import React, {useState} from "react";
import {RouteUrl} from "../Routes/routes";


let AuthContext = React.createContext<AuthContextType>(null!);

export interface AuthContextType {
    token: string
    user: {
        id: number
    };
}

//
// export const useAuth = () => React.useContext(AuthContext)
//
// export function AuthProvider({ children }: { children: React.ReactNode }) {
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


export function RequireAuth({children}: { children: JSX.Element }) {
    const {token, setToken} = useUserToken();
    let location = useLocation();

    console.log("RequireAuth")

    if (token) {
        console.log("!auth.user")

        return <Navigate to={RouteUrl.Signin} state={{from: location}} replace/>;
    }

    return children;
}


export function useUserToken() {
    const storageKeyName = "token";
    const getToken = () => {
        return sessionStorage.getItem(storageKeyName)
    }
    const [token, setToken] = useState(getToken());
    const repository = sessionRepository();


    const saveToken = (userToken: string) => {
        if (!userToken) {
            repository.remove()
        } else {
            repository.add(userToken)
        }

        setToken(userToken);
    };

    return {
        token,
        setToken: saveToken,
    }
}

interface SessionRepository {
    add(token: string): void

    remove(): void
}

export function sessionRepository(): SessionRepository {
    const storageKeyName = "token";

    return {
        add(token: string): void {
            return sessionStorage.setItem(storageKeyName, token);
        },
        remove() {
            return sessionStorage.removeItem(storageKeyName);
        }
    }
}

