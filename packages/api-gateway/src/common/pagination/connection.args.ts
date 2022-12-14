import {
  ConnectionArguments,
  ConnectionCursor,
  fromGlobalId,
} from "graphql-relay";

type PagingMeta =
  | { pagingType: "forward"; after?: string; first: number }
  | { pagingType: "backward"; before?: string; last: number }
  | { pagingType: "none" };

function checkPagingSanity(args: ConnectionArgs): PagingMeta {
  const { first = 0, last = 0, after, before } = args;

  const isForwardPaging = !!first || !!after;
  const isBackwardPaging = !!last || !!before;
  if (isForwardPaging && isBackwardPaging) {
    throw new Error("Relay pagination cannot be forwards AND backwards!");
  }
  if ((isForwardPaging && before) || (isBackwardPaging && after)) {
    throw new Error("Paging must use either first/after or last/before!");
  }
  if ((isForwardPaging && first < 0) || (isBackwardPaging && last < 0)) {
    throw new Error("Paging limit must be positive!");
  }
  if (last && !before) {
    throw new Error("When paging backwards, a 'before' argument is required!");
  }

  // eslint-disable-next-line no-nested-ternary
  return isForwardPaging
    ? { pagingType: "forward", after, first }
    : isBackwardPaging
    ? { pagingType: "backward", before, last }
    : { pagingType: "none" };
}

const getId = (cursor: ConnectionCursor) =>
  parseInt(fromGlobalId(cursor).id, 10);
const nextId = (cursor: ConnectionCursor) => getId(cursor) + 1;

function getPagingParameters(args: ConnectionArgs) {
  const meta = checkPagingSanity(args);

  switch (meta.pagingType) {
    case "forward": {
      return {
        limit: meta.first,
        offset: meta.after ? nextId(meta.after) : 0,
      };
    }
    case "backward": {
      const { last, before } = meta;
      let limit = last;
      let offset = getId(before!) - last;

      if (offset < 0) {
        limit = Math.max(last + offset, 0);
        offset = 0;
      }

      return { offset, limit };
    }
    default:
      return {};
  }
}

export default class ConnectionArgs implements ConnectionArguments {
  public before?: ConnectionCursor;

  public after?: ConnectionCursor;

  public first?: number;

  public last?: number;

  pagingParams() {
    return getPagingParameters(this);
  }
}
