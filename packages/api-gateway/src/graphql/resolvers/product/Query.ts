import ProductService from "src/adapters/product.adabter";
import {
  Product,
  ProductResponse,
  ProductsInput,
  ProductsPaginationResponse,
} from "src/graphql/types/products.types";

import dummyProduct from "./dummy-product";

export default {
  product: async (parent, productInput, context): Promise<any> => {
    console.log(productInput);
    let product = await ProductService.getProductDetails(productInput);
    console.log(product);
    return dummyProduct;
  },
  products: (
    parent: any,
    { input, pagination }: ProductsInput,
    context: any
  ): Promise<ProductsPaginationResponse> => {
    return new Promise<ProductsPaginationResponse>((resolve) => {
      console.log("parent", parent);
      console.log("input", input);
      console.log("pagination", pagination);
      console.log("type", new ProductsPaginationResponse());

      resolve([{ id: "id", productName: "product name" }]);
    });
  },
  productAssets: (parent: any, productAssetsInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
  category: (parent, categoryInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
  categories: (parent, categoriesInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
  subcategories: (parent, subcategoriesInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
  categoryProducts: (parent, categoryProductsInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
  categoryAssets: (parent: any, categoryAssetsInput, context): Promise<any> => {
    return new Promise((resolve) => {
      resolve(null);
    });
  },
};
