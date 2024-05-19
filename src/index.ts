import Logger from "loglevel";

Logger.enableAll();

function main() {
  const logger = Logger.getLogger("main");
  logger.info("Hello, world!");
}

main();
