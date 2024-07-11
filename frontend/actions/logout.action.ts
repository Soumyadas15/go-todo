"use server"

import { DEFAULT_LOGOUT_REDIRECT } from "@/routes";
import { redirect } from "next/dist/server/api-utils";
import { cookies } from "next/headers";
import { NextResponse } from "next/server";


export const signOut = async () => {
    try{

        const deletedCookie = cookies().delete('accessToken');

        return {
            success: true,
            redirectTo: DEFAULT_LOGOUT_REDIRECT,
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