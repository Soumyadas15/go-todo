"use client"

import { 
    Avatar, 
    AvatarFallback, 
    AvatarImage 
} from "@/components/ui/avatar"
import { User } from "@/types";
import { 
    DropdownMenu, 
    DropdownMenuItem, 
    DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { UserAvatar } from "@/components/user-avatar";
import { DropdownMenuContent } from "@radix-ui/react-dropdown-menu";
import { LogOut } from "lucide-react";
import { DEFAULT_LOGOUT_REDIRECT } from "@/routes";
import { useTransition } from "react";
import { useRouter } from "next/navigation";
import { signOut } from "@/actions/logout.action";
import toast from "react-hot-toast";

interface UserMenuProps{
    user: User;
}
export const UserMenu = ({
    user,
} : UserMenuProps) => {

    const router = useRouter();
    const [isPending, startTransition] = useTransition();

    const handleSignOut = async () => {
        startTransition(() => {
            signOut()
                .then((data : any) => {
                    if(data.redirect){
                        router.push(`${DEFAULT_LOGOUT_REDIRECT}`)
                    }
                })
                .catch((error) => {
                    toast.error("Unexpected error encountered")
                })
        })
    }


    return (

        <DropdownMenu>
            <DropdownMenuTrigger>
                <UserAvatar user={user}/>
            </DropdownMenuTrigger>

            <DropdownMenuContent 
                className="
                    w-36 
                    mt-2 
                    bg-white
                    dark:bg-neutral-900 
                    border-[1px] 
                    border-neutral-200 
                    dark:border-neutral-800
                    rounded-md p-1
                " 
                align="end"
            >
                <DropdownMenuItem 
                    className="cursor-pointer text-destructive flex items-center gap-2" 
                    onClick={handleSignOut}
                >
                    <LogOut size={20}/>
                    <p>Logout</p>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>

    )
}