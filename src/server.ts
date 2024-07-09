import http, { type IncomingMessage, type ServerResponse } from "node:http";
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
    const cost = (Date.now() - start).toString();
    console.log(`RPC Request ${req.method.name} took ${cost}ms`);
    return res;
  };
}

const routes = function (router: ConnectRouter) {
  router.service(Greet, { sayHello }, { interceptors: [logger()] });
};

const requestHandler = function (req: IncomingMessage, res: ServerResponse) {
  const start = Date.now();
  const query = new URL(req.url ?? "", "http://localhost:9000");
  const payload = new HelloRequest({
    language: query.searchParams.get("language") ?? "",
    name: query.searchParams.get("name") ?? "",
  });
  res.setHeader("Content-Type", "application/json");
  sayHello(payload).then((result) => {
    res.end(JSON.stringify(result));
    const cost = (Date.now() - start).toString();
    console.log(`HTTP Request ${req.url} took ${cost}ms`);
  });
};

http2.createServer(connectNodeAdapter({ routes })).listen(9090, () => {
  console.log("RPC Server on 9090");
});
const httpServer = http.createServer();
httpServer.on("request", requestHandler);
httpServer.listen(9000, () => {
  console.log("HTTP Server on 9000");
});
