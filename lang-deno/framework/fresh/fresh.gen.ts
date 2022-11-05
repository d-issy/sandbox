// DO NOT EDIT. This file is generated by fresh.
// This file SHOULD be checked into source version control.
// This file is automatically updated during development when running `dev.ts`.

import config from "./deno.json" assert { type: "json" };
import * as $0 from "./routes/[name].tsx";
import * as $1 from "./routes/about.jsx";
import * as $2 from "./routes/api/joke.ts";
import * as $3 from "./routes/api/random-uuid.ts";
import * as $4 from "./routes/countdown.tsx";
import * as $5 from "./routes/greet/[name].tsx";
import * as $6 from "./routes/index.tsx";
import * as $7 from "./routes/search.tsx";
import * as $$0 from "./islands/Countdown.tsx";
import * as $$1 from "./islands/Counter.tsx";

const manifest = {
  routes: {
    "./routes/[name].tsx": $0,
    "./routes/about.jsx": $1,
    "./routes/api/joke.ts": $2,
    "./routes/api/random-uuid.ts": $3,
    "./routes/countdown.tsx": $4,
    "./routes/greet/[name].tsx": $5,
    "./routes/index.tsx": $6,
    "./routes/search.tsx": $7,
  },
  islands: {
    "./islands/Countdown.tsx": $$0,
    "./islands/Counter.tsx": $$1,
  },
  baseUrl: import.meta.url,
  config,
};

export default manifest;
