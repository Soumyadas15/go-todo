"use client"

import { User } from "@/types";
import { useRouter } from "next/navigation";
import { Button } from "../ui/button";
import { ThemeToggle } from "../theme-toggle";
import { signOut } from "@/actions/logout.action";
import { useTransition } from "react";
import { DEFAULT_LOGOUT_REDIRECT } from "@/routes";
import { UserMenu } from "@/components/navbar/user-menu";
import { useModal } from "@/hooks/use-modal-store";

interface NavbarProps{
    user: User;
}
export const Navbar = ({
    user,
} : NavbarProps) => {

    const router = useRouter();
    const [isPending, startTransition] = useTransition();
    const { onOpen } = useModal();

    const addTodoClick = () => {
        if(!user){
            router.push('/auth/login')
            return
        }
        else {
            onOpen('todoModal')
        }
    }

    return (
        <nav className="
            fixed h-16 w-full 
            px-4 md:px-14 flex
            items-center justify-between
            bg-background
            transition
            duration-300
        ">
            <h1 
                className="text-2xl font-bold text-blue-500 cursor-pointer"
                onClick={() => router.push('/home')}
            >
                Todo.io
            </h1>
            <div className="flex items-center gap-4">
                <Button onClick={addTodoClick}>
                    Add todo
                </Button>
                <ThemeToggle/>
                {user && (
                    <UserMenu user={user}/>
                )}
            </div>
        </nav>
    )
}