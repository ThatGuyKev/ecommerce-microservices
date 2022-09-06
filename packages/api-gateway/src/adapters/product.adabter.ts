import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { PRODUCT_SERVICE_URI } from "config/default";

import { ProtoGrpcType } from "./_proto/product";
import { GetCategoriesByIdsReq } from "./_proto/product/GetCategoriesByIdsReq";
import { GetCategoryDetailsReq } from "./_proto/product/GetCategoryDetailsReq";
import { GetProductAssetByIdReq } from "./_proto/product/GetProductAssetByIdReq";
import { GetProductAssetsByProductIdReq } from "./_proto/product/GetProductAssetsByProductIdReq";
import { GetProductCategoriesReq } from "./_proto/product/GetProductCategoriesReq";
import { GetProductDetailsReq } from "./_proto/product/GetProductDetailsReq";
import { GetProductDetailsRes } from "./_proto/product/GetProductDetailsRes";
import { GetProductsByIdsReq } from "./_proto/product/GetProductsByIdsReq";
import { GetProductVariantsReq } from "./_proto/product/GetProductVariantsReq";

const packageDefinition = protoLoader.loadSync(
  `${__dirname}/_proto/product.proto`
);
const { product } = grpc.loadPackageDefinition(
  packageDefinition
) as unknown as ProtoGrpcType;

const client = new product.ProductService(
  PRODUCT_SERVICE_URI,
  grpc.credentials.createInsecure()
);
export default class ProductService {
  static async getProductDetails(input): Promise<GetProductDetailsRes> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductDetailsReq;
      console.log(request);
      client.getProductDetails(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getCategoryDetails({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetCategoryDetailsReq;
      client.getCategoryDetails(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getProductsByIds({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductsByIdsReq;
      client.getProductsByIds(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getCategoriesByIds({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetCategoriesByIdsReq;
      client.getCategoriesByIds(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getProductVariants({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductVariantsReq;
      client.getProductVariants(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getProductAssetsByProductId({
    input,
    pagination,
  }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductAssetsByProductIdReq;
      client.getProductAssetsByProductId(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getProductAssetById({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductAssetByIdReq;
      client.getProductAssetById(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });

    // depending on resp pass it back to client
  }
  static async getProductCategories({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = input as GetProductCategoriesReq;
      client.getProductCategories(request, (err, resp) => {
        if (err) {
          reject(err.message);
        }
        return resolve(resp);
      });
    });
  }
  static async getCategoryAssetsByCategoryId({
    input,
    pagination,
  }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = {};
      //   client.getCategoryAssetsByCategoryId(request, (err, resp) => {
      //     if (err) {
      //       reject(err.message);
      //     }
      //     return resolve(resp);
      //   });
    });

    // depending on resp pass it back to client
  }
  static async getCategoryAssetById({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = {};
      //   client.getCategoryAssetById(request, (err, resp) => {
      //     if (err) {
      //       reject(err.message);
      //     }
      //     return resolve(resp);
      //   });
    });

    // depending on resp pass it back to client
  }

  static async getCategoryProducts({ input, pagination }): Promise<any> {
    return new Promise((resolve, reject) => {
      let request = {};
      //   client.getCategoryProducts(request, (err, resp) => {
      //     if (err) {
      //       reject(err.message);
      //     }
      //     return resolve(resp);
      //   });
    });

    // depending on resp pass it back to client
  }
}
