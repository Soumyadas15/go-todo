"use server"

import { getCurrentUser } from "@/lib/get-current-user";
import axios from "axios";


interface Params {
  nextState?: string;
  sortBy?: string;
}

export const getTodos = async (params: Params = {}) => {

  try {

    const currentUser = await getCurrentUser();

    const { nextState, sortBy } = params;

    if (!currentUser) {
      return {
        error: "User not found",
      };
    }

    let apiUrl = `${process.env.BACKEND_URL}/api/todo/${currentUser.id}`;

    const pageParams = new URLSearchParams();

    if (nextState) {
        pageParams.append('pageState', nextState);
    }

    if (sortBy) {
        pageParams.append('sortBy', sortBy);
    }

    if (pageParams.toString()) {
        apiUrl += `?${pageParams.toString()}`;
    }

    console.log(apiUrl);

    const response = await axios.get(apiUrl);
    const { todos, nextPageState } = response.data;

    return {
      todos: todos,
      nextPageState: nextPageState,
    };


  } catch (error: any) {
    console.error('Error:', error);
    return {
      error: "An error occurred while processing your request.",
    };
  }
};
