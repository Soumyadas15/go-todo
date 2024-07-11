import { Todo } from "@/types";
import { 
    Card, 
    CardContent, 
    CardDescription, 
    CardFooter, 
    CardHeader,
    CardTitle
} from "@/components/ui/card";
import { CheckCheckIcon, Ellipsis, Pencil, Trash2 } from "lucide-react";
import { 
    DropdownMenu, 
    DropdownMenuContent, 
    DropdownMenuItem, 
    DropdownMenuSeparator, 
    DropdownMenuTrigger 
} from "@/components/ui/dropdown-menu";
import { useState, useTransition } from "react";
import { deleteTodo, markAsComplete } from "@/actions/todo.action";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useModal } from "@/hooks/use-modal-store";

interface TodoCardProps {
    todo: Todo;
}

export const TodoCard = ({ todo }: TodoCardProps) => {

    const [isPending, startTransition] = useTransition();
    const [loading, setLoading] = useState<boolean>(false);
    const router = useRouter();
    const { onOpen } = useModal();

    const handleDelete = () => {
        setLoading(true)
        startTransition(() => {
            deleteTodo(todo.id)
                .then((data: any) => {
                    toast.success('Todo deleted')
                })
                .catch((error) => {
                    toast.error('Something went wrong')
                })
                .finally(() => {
                    router.refresh();
                    setLoading(false);
                })
        })
    }


    const handleMarkAsComplete = async () => {
        setLoading(true)
        startTransition(() => {
            markAsComplete(todo.id)
                .then((data: any) => {
                    toast.success('Marked as complete')
                })
                .catch((error) => {
                    toast.error('Something went wrong')
                })
                .finally(() => {
                    router.refresh();
                    setLoading(false);
                })
        })
    }



    return (
        <Card className={`
            border-none
            ${todo.status === 'complete' ? 'bg-purple-200 dark:bg-purple-500/20' : 'bg-rose-200 dark:bg-rose-500/20'}
            transition duration-300
            ${loading ? 'opacity-50 cursor-not-allowed' : ''}
        `}>
            <CardHeader className="w-full flex flex-row items-center justify-between">
                <CardTitle>{todo.title}</CardTitle>
                <DropdownMenu>
                    <DropdownMenuTrigger className={`
                        ${loading ? 'cursor-not-allowed' : ''}
                    `}>
                        <Ellipsis />
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">

                        <DropdownMenuItem 
                            disabled={todo.status === 'complete'}
                            className="cursor-pointer gap-2"
                            onClick={() => {
                                onOpen('editTodoModal', todo.id, todo);
                            }}
                        >
                            <Pencil />
                            <p>Edit</p>
                        </DropdownMenuItem>

                        <DropdownMenuItem 
                            disabled={todo.status === 'complete'}
                            className="cursor-pointer gap-2"
                            onClick={handleMarkAsComplete}
                        >
                            <CheckCheckIcon />
                            <p>Mark as complete</p>
                        </DropdownMenuItem>

                        <DropdownMenuSeparator/>


                        <DropdownMenuItem 
                            className="text-red-500 cursor-pointer gap-2"
                            onClick={handleDelete}
                        >
                            <Trash2 />
                            <p>Delete</p>
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </CardHeader>
            <CardContent>
                <CardDescription>{todo.description}</CardDescription>
            </CardContent>
            <CardFooter>
                <div className={`
                    p-1
                    ${todo.status === 'complete' ? 'bg-purple-500' : 'bg-red-500'} 
                    rounded-full px-4 
                    text-white
                    flex items-center justify-center
                    `
                }>
                    <p>{todo.status}</p>
                </div>
            </CardFooter>
        </Card>
    );
};
