import { createPromiseClient } from "@connectrpc/connect";
import { createGrpcTransport } from "@connectrpc/connect-node";
import { Greet } from "../proto/greet/greet_connect.js";

const transport = createGrpcTransport({
  baseUrl: "http://localhost:8080",
  httpVersion: "2",
});

async function main() {
  const client = createPromiseClient(Greet, transport);
  const response = await client.sayHello({ name: "world" });
  console.log("response: ", response.message);
}

main();
