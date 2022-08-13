const { defineConfig } = require("@vue/cli-service");
const IS_PROD = process.env.NODE_ENV === "production";
module.exports = defineConfig({
  devServer: {
    host: "0.0.0.0",
    port: 9000,
  },
  css: {
    loaderOptions: {
      css: {
        modules: {
          //[name] 是文件名 [local]是具体的类名
          //类名模块化，防止重复
          //auto这样设置就可以防止vue文件中的style标签定义的样式丢失了
          // auto: /(\.module|base|style)\.scss$/i,
          auto: /\.scss/i,
          localIdentName: IS_PROD ? "[hash:8]" : "[local]-[hash:base64:6]",
        },
        esModule: false,
        sourceMap: !IS_PROD,
      },
    },
  },
  transpileDependencies: true,
  chainWebpack: (config) => {
    //启用ref sugar 语法糖
    config.module
      .rule("vue")
      .use("vue-loader")
      .tap((options) => {
        return {
          ...options,
          reactivityTransform: true,
        };
      });
    //css module不改变原来的名字
    if (!IS_PROD) {
      // config.module
      //   .rule("css")
      //   .use("style-loader")
      //   .end()
      //   .use("css-loader")
      //   .tap((options) => {
      //     console.log(options)
      //     return {
      //       ...options,
      //     }
      // })
    }
  },
});
