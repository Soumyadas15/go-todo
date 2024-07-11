"use client"

import { CreateTodoModal } from "@/components/modals/create-todo-modal"
import { EditTodoModal } from "@/components/modals/edit-todo-modal"

export const ModalProvider = () => {
    return (
        <>
            <EditTodoModal/>
            <CreateTodoModal/>  
        </>
    )
}