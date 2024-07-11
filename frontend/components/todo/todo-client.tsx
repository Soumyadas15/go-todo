"use client"

import { Todo } from "@/types";
import { TodoCard } from "@/components/todo/todo-card";
import { useModal } from "@/hooks/use-modal-store";
import { useEffect, useState } from "react";
import qs from "query-string";
import { Button } from "@/components/ui/button";
import { TodoSelect } from "./todo-select";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useRouter, useSearchParams } from "next/navigation";
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuTrigger } from "../ui/dropdown-menu";

interface TodoProps {
    todos: Todo[];
    nextState: string;
}

export const TodoClient = ({
    todos,
    nextState
} : TodoProps) => {

    const { onOpen } = useModal();
    const [loading, setLoading] = useState<boolean>(false);
    const [label, setLabel] = useState<string>('All todos');
    const router = useRouter();

    const searchParams = useSearchParams();
    const existingQuery = Object.fromEntries(
        searchParams.entries()
    )

    const handleNext = () => {
        const params = new URLSearchParams(existingQuery);

        if (nextState) {
        params.set("nextState", nextState);
        }

        const url = `/home?${params.toString()}`;
        router.push(url);
    };

    const handleSelect = (value: string) => {
        const params = new URLSearchParams(existingQuery);
        if (value === "") {
            setLabel("All todos");
            params.delete("sortBy");
        } else {
            setLabel(value.charAt(0).toUpperCase() + value.slice(1));
            params.set("sortBy", value);
        }
        const url = `/home?${params.toString()}`;
        router.push(url);
    };

    if (!todos) {
        return (
            <div className="h-full w-full flex items-center justify-center">
                <div className="flex flex-col items-center gap-2">
                    <h1 className="text-2xl font-bold">
                        Create your first todo
                    </h1>
                    <Button
                        onClick={() => {
                            onOpen('todoModal')
                        }}
                    >
                        Create todo
                    </Button>
                </div>
            </div>
        )
    }

    return (
        <div className="w-full pt-20 px-4 md:px-14 flex flex-col items-end gap-4">
            <DropdownMenu>
                <DropdownMenuTrigger className="w-56">
                    <Button className="w-full flex items-start" variant={'outline'}>
                        {label}
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-56">
                    <DropdownMenuGroup>
                        <DropdownMenuItem onClick={() => handleSelect('')}>
                            All tasks
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => handleSelect('pending')}>
                            Pending
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => handleSelect('complete')}>
                            Complete
                        </DropdownMenuItem>
                    </DropdownMenuGroup>
                </DropdownMenuContent>
            </DropdownMenu>
            <div className="
                grid 
                grid-cols-1 sm:grid-cols-1 
                md:grid-cols-2 lg:grid-cols-4 
                xl:grid-cols-4 2xl:grid-cols-4 
                gap-2
            ">
                {todos.map((todo: Todo, index: number) => (
                    <div key={index}>
                        <TodoCard 
                            key={index} 
                            todo={todo}
                        />
                    </div>
                ))}
            </div>
            <Button 
                variant="outline"
                onClick={handleNext}
                disabled={loading}
            >
                Next Page
            </Button>
        </div>
    )
}