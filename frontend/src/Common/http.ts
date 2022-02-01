import {HeaderKey} from "../Auth/constants";
import {sessionRepository, useUserToken} from "../Auth/RequireAuth";
import {RouteUrl} from "../Routes/routes";

const repository = sessionRepository();

export async function httpAuthApiJson(endpoint: string, token: string, params?: RequestInit): Promise<Response> {
    const headers = new Headers(params?.headers)

    headers.set(HeaderKey.AuthToken, token)

    console.log({headers})

    const res = await httpApiJson(endpoint, {
        ...params,
        headers,
    });
    if (res.status === 401) {
        repository.remove();
        window.location.href = "/"
    }

    return res
}

export function httpApiJson(endpoint: string, params?: RequestInit): Promise<Response> {

    const headers = new Headers(params?.headers)

    console.log({headers})

    headers.set('Content-Type', 'application/json')

    return fetch(`${process.env.REACT_APP_API_URL || ""}${endpoint}`, {
        ...params,
        headers,
    })
}