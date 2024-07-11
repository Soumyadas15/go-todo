import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import { cookies, headers } from "next/headers";
import {
    DEFAULT_LOGIN_REDIRECT,
    DEFAULT_LOGOUT_REDIRECT,
    authRoutes,
} from '@/routes'
import { getTodos } from './actions/get-todos.action';

export function middleware(request: NextRequest) {

    const { nextUrl } = request;

    let user = cookies().get('accessToken')?.value ?? "";
    
    const isAuthRoute = authRoutes.includes(nextUrl.pathname);

    let isLoggedIn = false;

    if (user) {
        isLoggedIn = true;
    }

    console.log(isLoggedIn);
    

    if(isAuthRoute){
        if(isLoggedIn){
            return Response.redirect(new URL(DEFAULT_LOGIN_REDIRECT, nextUrl));
        }
        return null;
    }

    if(!isLoggedIn){
        return Response.redirect(new URL(DEFAULT_LOGOUT_REDIRECT, nextUrl));
    }

    if (nextUrl.pathname === '/'){
        return Response.redirect(new URL(DEFAULT_LOGIN_REDIRECT, nextUrl));
    }
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
}