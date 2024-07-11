import { getTodos } from '@/actions/get-todos.action';
import { TodoClient } from '@/components/todo/todo-client';
import axios from 'axios';
import Cookies from 'js-cookie';
import { cookies } from "next/headers"
// import { base64Encode } from 'js-base64';



const HomePage = async ({
    params,
    searchParams,
} : {
    params: { slug: string };
    searchParams?: { [key: string]: string | string[] | undefined };
}) => {

    cookies().getAll().map((cookie : any) => {
        console.log(cookie)
    })

    //@ts-ignore
    const { todos, nextPageState } = await getTodos(searchParams);


    return ( 
        <div className='bg-background transition duration-300 h-full w-full'>
            <TodoClient
                todos={todos}
                nextState={nextPageState}
            />
        </div>
     );
}
 
export default HomePage;