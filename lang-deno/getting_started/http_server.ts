import { serve } from "https://deno.land/std@0.162.0/http/server.ts";

function handler(req: Request): Response {
  console.log("Method: ", req.method);

  const url = new URL(req.url);
  console.log("Path: ", url.pathname);
  console.log("Query Parameters: ", url.searchParams);

  console.log("Headers: ", req.headers);

  if (req.body) {
    const body = req.text();
    console.log("Body: ", body);
  }

  return new Response("Hello, World!");
}

serve(handler, { port: 4242 });
