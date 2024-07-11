import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Cookies from 'js-cookie';
import { getCurrentUser } from "@/lib/get-current-user";
import { Navbar } from "@/components/navbar/navbar";
import { ThemeProvider } from "@/providers/theme-provider";
import { ModalProvider } from "@/providers/modal-provider";
import { Toaster } from "react-hot-toast";
import { CookiesProvider } from 'next-client-cookies/server';

export default async function DefaultLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const currentUser = await getCurrentUser();
  console.log(currentUser.id)

  return (
    <html lang="en">
      <body className="h-screen">
        <CookiesProvider>
          <ThemeProvider
              attribute="class"
              defaultTheme="system"
          >
            <Toaster/>
              <Navbar user={currentUser}/>
              <ModalProvider/>
              {children}
          </ThemeProvider>
        </CookiesProvider>
      </body>
    </html>
  );
}
