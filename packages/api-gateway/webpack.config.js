const path = require("path");
const webpackNodeExternals = require("webpack-node-externals");
const TsconfigPaths = require("tsconfig-paths-webpack-plugin");
module.exports = {
  entry: "./src/index.ts",
  mode: "production",
  externals: [webpackNodeExternals()],
  module: {
    rules: [
      {
        test: /\.ts?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".js"],
    plugins: [
      new TsconfigPaths({
        configFile: "./tsconfig.json",
      }),
    ],
  },
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "dist"),
  },
  target: "node",
};
