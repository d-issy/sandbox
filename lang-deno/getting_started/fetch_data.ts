import { readableStreamFromReader, writableStreamFromWriter } from "https://deno.land/std@0.162.0/streams/mod.ts";

/**
 * Output: JSON Data
 */
const jsonResponse = await fetch("https://api.github.com/users/denoland");
const jsonData = await jsonResponse.json();
console.log(jsonData);

/**
 * Output: HTML Data
 */
const textResponse = await fetch("https://deno.land/");
const textData = await textResponse.text();
console.log(textData);

/**
 * Output: Error Message
 */
try {
  await fetch("https:/does.not.exist/");
} catch (error) {
  console.log(error);
}

/**
 * Receiving a file
 */
const fileResponse = await fetch("https://deno.land/logo.svg");
if (fileResponse.body) {
  const file = await Deno.open("./data/logo.svg", {
    write: true,
    create: true,
  });
  const writableStream = writableStreamFromWriter(file);
  await fileResponse.body.pipeTo(writableStream);
}

/**
 * Sending a file
 */
const file = await Deno.open("./data/logo.svg", { read: true });
const readableStream = readableStreamFromReader(file);

await fetch("https://example.com/", {
  method: "POST",
  body: readableStream,
});
