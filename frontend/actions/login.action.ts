"use server"

import { LoginSchema } from "@/schemas"
import * as z from "zod";
import axios from "axios";
import { cookies } from "next/headers";
import { DEFAULT_LOGIN_REDIRECT } from "@/routes";

export const login = async (values: z.infer<typeof LoginSchema>) => {
    try {
        const validatedFields = LoginSchema.safeParse(values);

        if (!validatedFields.success) {
            return {
                error: "Invalid fields!",
            };
        }

        const { email, password } = validatedFields.data;
        const data = { email, password };

        const response = await axios.post(`${process.env.BACKEND_URL}/api/auth/login`, data,
            {
                headers: {
                    'Content-Type': 'application/json',
                },
                withCredentials: true,
            })


        const { user, token } = response.data;

        cookies().set("userToken", JSON.stringify(user), {
            httpOnly: true,
            maxAge: 24 * 60 * 60,
            sameSite: "strict"
        });

        cookies().set("jwtToken", token, {
            httpOnly: true,
            maxAge: 24 * 60 * 60,
            sameSite: "strict"
        });

        return {
            success: "OK",
            redirect: DEFAULT_LOGIN_REDIRECT,
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

