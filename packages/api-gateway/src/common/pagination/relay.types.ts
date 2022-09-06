import * as Relay from "graphql-relay";
import PageData from "./page-data";

interface Type<T = any> extends Function {
  new (...args: any[]): T;
}
const typeMap = {};
export default function relayTypes<T>(type: Type<T>): any {
  const { name } = type;
  if (typeMap[`${name}`]) return typeMap[`${name}`];

  class Connection implements Relay.Connection<T> {
    public name = `${name}Connection`;

    public edges!: Relay.Edge<T>[];

    public pageInfo!: Relay.PageInfo;
  }

  abstract class Page {
    public name = `${name}Page`;

    public page!: Connection;

    public pageData!: PageData;
  }

  // class Edge implements Relay.Edge<T> {
  //   public name = `${name}Edge`;

  //   public cursor!: Relay.ConnectionCursor;

  //   public node!: T;
  // }

  // class PageInfo implements Relay.PageInfo {
  //   public startCursor!: Relay.ConnectionCursor;

  //   public endCursor!: Relay.ConnectionCursor;

  //   public hasPreviousPage!: boolean;

  //   public hasNextPage!: boolean;
  // }

  typeMap[`${name}`] = Page;
  return typeMap[`${name}`];
}
