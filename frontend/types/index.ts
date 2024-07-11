export type User = {
    id: string; 
    username: string;
    email: string;
    password: string;
    todos?: Todo[];
    createdAt: string;
};

export type Todo = {
    id: string; 
    title: string;
    description: string;
    status: string;
    userId: string;
    backgroundColor: string;
    createdAt: string;
    updatedAt: string;
};
