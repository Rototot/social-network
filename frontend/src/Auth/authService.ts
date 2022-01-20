
interface SinginInputInterface{
    email: string
    password: string
}

interface SignupInputInterface{
    firstName: string
    lastName: string
    city: string
    interests: string
    email: string
    password: string
}

export async function SignIn(payload: SinginInputInterface) {
    const res = await fetch(`${process.env.REACT_APP_API_URL || ""}/api/signin`, {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: {'Content-Type': 'application/json'},
    })

    // eslint-disable-next-line no-console
    console.log({
        data: payload,
    });

    if (res.status === 200) {
        const resBody = await res.json()

        // eslint-disable-next-line no-console
        console.log({
            resBody,
        });
    }
}

export function SignUp(payload: SignupInputInterface) {

}

export function SignOut() {

}