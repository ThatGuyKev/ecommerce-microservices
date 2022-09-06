import { gql } from "apollo-server";

import ProductSchema from "./product.schema";
export const schema = gql`
  ${ProductSchema}
`;
