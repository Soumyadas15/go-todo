"use server"

import { RegisterSchema } from "@/schemas"
import * as z from "zod";
import axios from "axios"
import { login } from "./login.action";


export const register = async (values: z.infer<typeof RegisterSchema>) => {
    try {
        const validatedFields = RegisterSchema.safeParse(values);

        if (!validatedFields.success) {
            return {
                error: "Invalid fields!",
            };
        }

        const { username, email, password } = validatedFields.data;
        const data = { username, email, password };

        const response = await axios.post(
            `${process.env.BACKEND_URL}/api/auth/register`,
            data,
            {
                headers: {
                    'Content-Type': 'application/json',
                },
            }
        );
        

        return {
            success: "Success",
            redirect: true,
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