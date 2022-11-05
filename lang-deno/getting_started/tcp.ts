import { copy } from "https://deno.land/std@0.162.0/streams/conversion.ts";

const hostname = "0.0.0.0";
const port = 8080;
const listener = Deno.listen({ hostname, port });

console.log(`listening on ${hostname}:${port}`);

for await (const conn of listener) {
  copy(conn, conn);
}
