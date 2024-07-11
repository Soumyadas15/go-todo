"use client"

import { 
    Avatar, 
    AvatarFallback, 
    AvatarImage 
} from "@/components/ui/avatar"
import { User } from "@/types";

interface UserAvatarProps{
    user: User;
}
export const UserAvatar = ({
    user,
} : UserAvatarProps) => {

    return (
        <Avatar>
            <AvatarImage src={"https://github.com/shadcn.png"} alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
        </Avatar>
    )
}