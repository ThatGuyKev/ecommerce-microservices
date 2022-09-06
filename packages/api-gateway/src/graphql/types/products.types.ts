import { Request, Response } from "express";
import ConnectionArgs from "src/common/pagination/connection.args";
import relayTypes from "src/common/pagination/relay.types";

export type ProductsInput = {
  input: any;
  pagination: ConnectionArgs;
};
export interface ResolverContext {
  req: Request;
  res: Response;
}
export class Product {
  productName: string;
}

export class ProductResponse extends Product {}
export class ProductsPaginationResponse extends relayTypes<Product>(Product) {}
