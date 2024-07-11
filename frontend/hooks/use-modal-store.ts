import { create } from "zustand";
import { Todo } from "@/types";

export type ModalType = "todoModal" | "deleteModal" | "editTodoModal";

interface ModalStore {
    type: ModalType | null;
    isOpen: boolean;
    todoId: string | null;
    todo: Todo | null;
    onOpen: (type: ModalType, todoId?: string, todo?: Todo) => void;
    onClose: () => void;
    clearTodo: () => void;
}

export const useModal = create<ModalStore>((set) => ({
    type: null,
    isOpen: false,
    todoId: null,
    todo: null,
    onOpen: (type, todoId, todo) => set({ type, isOpen: true, todoId, todo }),
    onClose: () => set({ type: null, isOpen: false, todoId: null, todo: null }),
    clearTodo: () => set({ todo: null }),
}));
