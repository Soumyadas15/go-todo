"use server"

import { cookies } from "next/headers";

// setHuhCookie.action.js
export default async function setHuhCookie(huhValue : string) {
  return cookies().set("huh", JSON.stringify(huhValue), {
    httpOnly: true,
    maxAge: 24 * 60 * 60,
    sameSite: "strict",
  });
}
