// Global write permission.
const desc1 = { name: "write" } as const;
console.log(await Deno.permissions.request(desc1));

// Write permission to `$PWD/foo/bar`.
const desc2 = { name: "write", path: "foo/bar" } as const;
console.log(await Deno.permissions.request(desc2));

// Global net permission.
const desc3 = { name: "net" } as const;
console.log(await Deno.permissions.request(desc3));

// Net permission to 127.0.0.1:8000.
const desc4 = { name: "net", host: "127.0.0.1:8000" } as const;
console.log(await Deno.permissions.request(desc4));

// High-resolution time permission.
const desc5 = { name: "hrtime" } as const;
console.log(await Deno.permissions.request(desc5));
