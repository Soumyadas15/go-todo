import { cookies } from "next/headers";


export const getCookie = async (name: string) => {
  return cookies().get(name)?.value ?? "";
};


export async function getCurrentUser() {
    const userToken = await getCookie("userToken");
    if(userToken){
        return JSON.parse(userToken);
    }
    else{
        return null;
    }
}