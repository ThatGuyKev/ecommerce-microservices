import { accessEnv } from "access-env";
export const PORT = accessEnv("PORT", 9000);
export const PRODUCT_SERVICE_URI = accessEnv(
  "PRODUCT_SERVICE_URI",
  "0.0.0.0:9101"
);
