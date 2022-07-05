const { merge } = require("webpack-merge");
const path = require("path");
const commonConfig = require("./webpack.common");
const { SERVER_HOST, SERVER_PORT, ROOT_PATH } = require("../constant");
const webpack = require("webpack");

/**
 * @type {(import 'webpack').Configuration}
 */
const DevConfig = {
  target: "web",
  mode: "development",
  devtool: "eval-source-map",
  devServer: {
    host: SERVER_HOST,
    port: SERVER_PORT,
    compress: true, // gzip压缩
    open: true, // 自动打开默认浏览器
    hot: true, // 启用服务热替换配置
    client: {
      logging: "warn", // warn以上的信息，才会打印
      overlay: true, // 当出现编译错误或警告时，在浏览器中显示全屏覆盖
    },
    // 解决路由跳转404问题
    historyApiFallback: true,
  },
  // optimization:{

  // }
};

module.exports = merge(commonConfig, DevConfig);
