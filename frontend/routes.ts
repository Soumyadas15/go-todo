/**
 * An array of routes that are used for authentication
 * These routes will redirect logged in users to /settings
 * @type {string[]}
 */


export const authRoutes = [
  "/auth/login",
  "/auth/register",
];



export const DEFAULT_LOGIN_REDIRECT = "/home";
export const DEFAULT_LOGOUT_REDIRECT = "/auth/login";