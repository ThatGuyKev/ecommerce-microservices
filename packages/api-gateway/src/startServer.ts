import { ApolloServer } from "apollo-server";
import express from "express";
import { PORT } from "config/default";

import { schema } from "src/graphql/schema";
import resolvers from "src/graphql/resolvers";

import formatGraphQLErrors from "./formatGraphQLErrors";

import { ApolloServerPluginLandingPageGraphQLPlayground } from "apollo-server-core";
(async function startServer() {
  const apolloServer = new ApolloServer({
    context: (a) => a,
    formatError: formatGraphQLErrors,
    resolvers,
    typeDefs: schema,
    plugins: [ApolloServerPluginLandingPageGraphQLPlayground],
  });

  const app = express();

  // apolloServer.applyMiddleware({ app, cors: false, path: "/graphql" });
  let info = await apolloServer.listen(PORT, "0.0.0.0");
  console.info(`API gateway listening on: ${info.port}`);
})();
