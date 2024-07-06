import http2 from "node:http2";
import { connectNodeAdapter } from "@connectrpc/connect-node";
import { ConnectRouter, Interceptor } from "@connectrpc/connect";
import { Greet } from "../proto/greet/greet_connect.js";
import { HelloRequest, HelloResponse } from "../proto/greet/greet_pb.js";

const sayHello = async function (req: HelloRequest): Promise<HelloResponse> {
  const res = new HelloResponse();
  let prefix = "Hello";
  switch (req.language.toLowerCase()) {
    case "es":
      prefix = "Hola";
      break;
    case "fr":
      prefix = "Bonjour";
      break;
    case "de":
      prefix = "Hallo";
      break;
    default:
      break;
  }
  res.message = `${prefix}, ${req.name}`;
  return res;
};

function logger(): Interceptor {
  return (next) => async (req) => {
    const start = Date.now();
    const res = await next(req);
    console.log(
      `Request ${req.method.name} took ${(Date.now() - start).toString()}ms`
    );
    return res;
  };
}

const routes = function (router: ConnectRouter) {
  router.service(Greet, { sayHello }, { interceptors: [logger()] });
};

async function main() {
  http2.createServer(connectNodeAdapter({ routes })).listen(9000);
}

main();
