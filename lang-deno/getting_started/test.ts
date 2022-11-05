import { assertEquals, assertExists } from "https://deno.land/std@0.162.0/testing/asserts.ts";
import { delay } from "https://deno.land/std@0.162.0/async/delay.ts";

Deno.test("url test", () => {
  const url = new URL("./foo.js", "https://deno.land/");
  assertEquals(url.href, "https://deno.land/foo.js");
});

Deno.test("hello world #1", () => {
  const x = 1 + 2;
  assertEquals(x, 3);
});

Deno.test(function helloWorld3() {
  const x = 1 + 2;
  assertEquals(x, 3);
});

Deno.test({
  name: "hello world #2",
  fn: () => {
    const x = 1 + 2;
    assertEquals(x, 3);
  },
});

Deno.test("async hello world", async () => {
  const x = 1 + 2;
  await delay(100);
  if (x !== 3) {
    throw Error("x should be equal to 3");
  }
});

Deno.test("exists", () => {
  assertExists("Denosaurus");
  assertExists(false);
  assertExists(0);
});
