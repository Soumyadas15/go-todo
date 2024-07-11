"use client"


import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form"
import { Modal } from "@/components/modals/modal"
import { useForm } from "react-hook-form"
import { Input } from "@/components/ui/input"
import { useCallback, useEffect, useState, useTransition } from "react"
import { useModal } from "@/hooks/use-modal-store"
import { useRouter } from "next/navigation"
import toast from "react-hot-toast"
import { Textarea } from "../ui/textarea"
import { updateTodo } from "@/actions/todo.action"


export const EditTodoModal = () => {

    const { isOpen, onClose, type, todo, clearTodo } = useModal();

    const router = useRouter();
    const form = useForm({
        defaultValues: {
            //@ts-ignore
          title: todo?.title,
          description: todo?.description,
          todoId: todo?.id,
        },
    });


    const [loading, setLoading] = useState<boolean>(false);
    const [isPending, startTransition] = useTransition();



    const isModalOpen = isOpen && type === "editTodoModal"


    useEffect(() => {
        if (todo) {
            form.setValue('title', todo.title || '');
            form.setValue('description', todo.description || '');
            form.setValue('todoId', todo.id || '');
            // form.setValue('status', task.status);
        }
    }, [todo]);



    const onSubmit = async (values: any) => {
        setLoading(true);
        startTransition(() => {
            updateTodo(values)
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

        form.reset();
        // clearTodo();
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
                title="Edit todo"
                onClose={handleClose}
                onSubmit={form.handleSubmit(onSubmit)}
                actionLabel="Update"
                secondaryActionLabel="Cancel"
                secondaryAction={handleClose}
                isOpen={isModalOpen}
                body={bodyContent}
                disabled={loading}
            />
        </div>
    )
}