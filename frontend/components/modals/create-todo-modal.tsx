"use client"


import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Modal } from "@/components/modals/modal"
import { FieldValues, SubmitHandler, useForm } from "react-hook-form"
import { Input } from "@/components/ui/input"
import { useCallback, useEffect, useMemo, useState, useTransition } from "react"
import { useModal } from "@/hooks/use-modal-store"
import { AnimatedDiv } from "@/components/ui/animated-div"
import { useRouter } from "next/navigation"
import { TodoSchema } from "@/schemas";
import * as z from 'zod';
import axios from "axios"
import toast from "react-hot-toast"
import { Textarea } from "../ui/textarea"
import { createTodo } from "@/actions/todo.action"


export const CreateTodoModal = () => {

    const router = useRouter();
    const form = useForm<z.infer<typeof TodoSchema>>({
        defaultValues: {
          title: "",
          description: "",
        },
    });

  
    const { isOpen, onClose, type } = useModal();

    const [loading, setLoading] = useState<boolean>(false);
    const [isPending, startTransition] = useTransition();



    const isModalOpen = isOpen && type === "todoModal"



    const onSubmit = async (values: z.infer<typeof TodoSchema>) => {
        setLoading(true);
        startTransition(() => {
            createTodo(values)
                .then((data : any) => {
                    toast.success('Success')
                })
                .catch((error) => {
                    toast.error('Something went wrong')
                })
                .finally(() => {
                    router.refresh();
                    handleClose();
                    setLoading(false);
                })
        })
    }
    
    


    const handleClose = useCallback(() => {

        //@ts-ignore
        form.setValue('title', '');
        form.setValue('description', '');

        form.reset()
        onClose();
    }, []);


    let bodyContent = (
        <div className="flex flex-col gap-4">
            <Form {...form}>
                <FormField
                    control={form.control}
                    name="title"
                    render={({ field }) => (
                        <FormItem>
                            <FormControl>
                                <Input 
                                    className="focus-visible:ring-transparent focus:ring-0" 
                                    placeholder="Title" 
                                    {...field} 
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                        <FormItem>
                        <FormControl>
                            <Textarea 
                                className="focus-visible:ring-transparent focus:ring-0 h-[15rem]" 
                                placeholder="Description" 
                                {...field} 
                            />
                        </FormControl>
                        <FormMessage />
                        </FormItem>
                    )}
                />
            </Form>
        </div>
    )


    return (
        <div>
            <Modal
                title="Create a new todo"
                onClose={handleClose}
                onSubmit={form.handleSubmit(onSubmit)}
                actionLabel="Create"
                secondaryActionLabel="Cancel"
                secondaryAction={handleClose}
                isOpen={isModalOpen}
                body={bodyContent}
                disabled={loading}
            />
        </div>
    )
}