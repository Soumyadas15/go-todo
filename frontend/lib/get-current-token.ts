import { cookies } from "next/headers";
import { getCookie } from "./get-current-user";


export async function getCurrentToken() {
    const jwtToken = await getCookie("jwtToken");
    if(jwtToken){
        return jwtToken;
    }
    else{
        return null;
    }
}