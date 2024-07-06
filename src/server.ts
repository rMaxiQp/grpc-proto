import https from "node:https";
import { HelloRequest, HelloResponse } from "../proto/greet/greet_pb.js";

function greet(payload: HelloRequest): Promise<HelloResponse> {
  return new Promise((resolve, reject) => {
    const req = https.request(
      {
        host: "localhost",
        port: 8080,
        method: "POST",
      },
      (res) => {
        let data = "";
        res.on("data", (chunk) => {
          data += chunk;
        });
        res.on("end", () => {
          resolve(HelloResponse.fromJson(data));
        });
        res.on("error", reject);
      }
    );

    req.write(JSON.stringify(payload));
    req.end();
  });
}

async function main() {
  // const server = new Server();
  // server.addService(Greet, { sayHello });
  // client.makeUnaryRequest()

  const payload = new HelloRequest({ name: "world" });
  await greet(payload);
  console.log("Hello, world!");
}

main();
