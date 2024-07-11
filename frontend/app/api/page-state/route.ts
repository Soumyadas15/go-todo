import { cookies } from "next/headers";
import { NextResponse } from "next/server";
import { serialize } from "cookie";
import { NextApiRequest, NextApiResponse } from "next";

export async function POST(request: Request, res: NextApiResponse) {
  try {
    const { nextPageState } = await request.json();

    const cookieValue = 'some-value';

    const cookie = serialize('myCookie', cookieValue, {
      httpOnly: false,
      secure: process.env.NODE_ENV === 'production',
      maxAge: 60 * 60 * 24 * 7,
      sameSite: 'strict',
      path: '/',
    });

    return NextResponse.json({ message: 'Cookie set!' }, {
      headers: {
        'Set-Cookie': cookie,
      },
    });

  } catch (error) {
    console.error('Error setting cookie:', error);

    if (error instanceof SyntaxError) {
      return NextResponse.json({ message: 'Invalid JSON data' }, { status: 400 });
    } else if (error instanceof TypeError) {
      return NextResponse.json({ message: 'Missing nextPageState in request body' }, { status: 400 });
    } else {
      return NextResponse.json({ message: 'Internal Server Error' }, { status: 500 });
    }
  }
}
