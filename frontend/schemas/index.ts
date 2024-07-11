import * as z from "zod";

export const LoginSchema = z.object({
    email: z.string().email({
        message: "Email is required"
    }),
    password: z.string().min(1, {
        message: "Name is required"
    }),
})

export const RegisterSchema = z.object({
    username: z.string().min(1, {
        message: "Username is required"
    }),
    email: z.string().email({
        message: "Email is required"
    }),
    password: z.string().min(6, {
        message: "Minimum 6 characters required"
    })
})

export const TodoSchema = z.object({
    title: z.string().min(1, {
        message: "Title is required"
    }),
    description: z.string().email({
        message: "Description is required"
    }),
})