import { cookies } from "next/headers";


export const getCookie = async (name: string) => {
  return cookies().get(name)?.value ?? "";
};


export async function getCurrentUser() {
    const accessToken = await getCookie("accessToken");
    if(accessToken){
        return JSON.parse(accessToken);
    }
    else{
        return null;
    }
}