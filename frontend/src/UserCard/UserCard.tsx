import {useEffect, useState} from "react";

export default function UserCard(){
    // const [error, setError] = useState(null);
    // const [isLoaded, setIsLoaded] = useState(false);
    // const [userInfo, setUserInfo] = useState([]);



    // useEffect(() => {
    //     fetch("https://api.example.com/items")
    //         .then(res => res.json())
    //         .then(
    //             (result) => {
    //                 setIsLoaded(true);
    //                 setUserInfo(result);
    //             },
    //             // Примечание: важно обрабатывать ошибки именно здесь, а не в блоке catch(),
    //             // чтобы не перехватывать исключения из ошибок в самих компонентах.
    //             (error) => {
    //                 setIsLoaded(true);
    //                 setError(error);
    //             }
    //         )
    // }, [])
    //
    // if (error) {
    //     return <div>Ошибка: {error.message}</div>;
    // } else if (!isLoaded) {
    //     return <div>Загрузка...</div>;
    // }

    return (
        <>
            {/*<ul>*/}
            {/*    {items.map(item => (*/}
            {/*        <li key={item.id}>*/}
            {/*            {item.name} {item.price}*/}
            {/*        </li>*/}
            {/*    ))}*/}
            {/*</ul>*/}
        </>
    );
}