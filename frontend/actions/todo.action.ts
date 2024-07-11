"use server"

import { TodoSchema } from "@/schemas"
import * as z from "zod";
import axios from "axios";
import { getCurrentUser } from "@/lib/get-current-user";
import { cookies } from "next/headers";



export const createTodo = async (values: z.infer<typeof TodoSchema>) => {
    try {

        const { title, description } = values;
        const currentUser = await getCurrentUser();

        if(!currentUser){
            return {
                error: "User not found",
            };
        }

        const data = {
            title,
            description,
            userId: currentUser.id, 
        }; 

        const response = await axios.post(`${process.env.BACKEND_URL}/api/todo`, data,
            {
                headers: {
                    'Content-Type': 'application/json',
                },
            })

        const { todo } = response.data;

        return {
            success: "OK",
            todo: todo,
        };


    } catch (error : any) {
        if (error.response) {
            console.error('Error:', error.response.data);
            return {
                error: error.response.data,
            };
        } else if (error.request) {
            console.error('Error:', error.request);
            return {
                error: "No response received from server.",
            };
        } else {
            console.error('Error:', error.message);
            return {
                error: "An error occurred while processing your request.",
            };
        }
    }
}



export const updateTodo = async (values: any) => {
    try {

        const { title, description, todoId } = values;

        const currentUser = await getCurrentUser();

        if(!currentUser){
            return {
                error: "User not found",
            };
        }

        const data = {
            title,
            description,
            status: "pending",
            userId: currentUser.id
        }; 
        
        const apiUrl = `${process.env.BACKEND_URL}/api/todo/${todoId}`

        const response = await axios.put(apiUrl, data,
            {
                headers: {
                    'Content-Type': 'application/json',
                },
            })

        return {
            success: "OK",
        };


    } catch (error : any) {
        if (error.response) {
            console.error('Error:', error.response.data);
            return {
                error: error.response.data,
            };
        } else if (error.request) {
            console.error('Error:', error.request);
            return {
                error: "No response received from server.",
            };
        } else {
            console.error('Error:', error.message);
            return {
                error: "An error occurred while processing your request.",
            };
        }
    }
}

// export const getTodos = async (pageState : string | null) => {
//     try {

//         const currentUser = await getCurrentUser();

//         if(!currentUser){
//             return {
//                 error: "User not found",
//             };
//         }

//         let apiUrl = `${process.env.BACKEND_URL}/api/todo/${currentUser.id}`;
//         if (pageState) {
//             apiUrl += `?pageState=${pageState}`;
//         }


//         const response = await axios.get(apiUrl);
//         const { todos, nextPageState } = response.data;

//         return {
//             todos: todos,
//             nextPageState: nextPageState
//         };

//     } catch (error : any) {
//         if (error.response) {
//             console.error('Error:', error.response.data);
//             return {
//                 error: error.response.data,
//             };
//         } else if (error.request) {
//             console.error('Error:', error.request);
//             return {
//                 error: "No response received from server.",
//             };
//         } else {
//             console.error('Error:', error.message);
//             return {
//                 error: "An error occurred while processing your request.",
//             };
//         }
//     }
// }


export const deleteTodo = async (todoId : string) => {
    try {

        const currentUser = await getCurrentUser();
        const userId = currentUser.id;

        if(!currentUser){
            return {
                error: "User not found",
            };
        }


        const response = await axios.delete(`${process.env.BACKEND_URL}/api/todo/${todoId}`, {
            data: { userId }
        });

        return {
            success: 'Ok'
        }

    } catch (error : any) {
        if (error.response) {
            console.error('Error:', error.response.data);
            return {
                error: error.response.data,
            };
        } else if (error.request) {
            console.error('Error:', error.request);
            return {
                error: "No response received from server.",
            };
        } else {
            console.error('Error:', error.message);
            return {
                error: "An error occurred while processing your request.",
            };
        }
    }
}


export const markAsComplete = async (todoId : string) => {
    try{

        const currentUser = await getCurrentUser();
        const userId = currentUser.id;
        const data = { userId }

        console.log(data);
        console.log(todoId);

        if(!currentUser){
            return {
                error: "User not found",
            };
        }
        const apiUrl = `${process.env.BACKEND_URL}/api/todo/mark-complete/${todoId}`
        const response = await axios.put(apiUrl, data);

        return {
            success: 'Ok'
        }

    } catch (error : any) {
        if (error.response) {
            console.error('Error:', error.response.data);
            return {
                error: error.response.data,
            };
        } else if (error.request) {
            console.error('Error:', error.request);
            return {
                error: "No response received from server.",
            };
        } else {
            console.error('Error:', error.message);
            return {
                error: "An error occurred while processing your request.",
            };
        }
    }
}