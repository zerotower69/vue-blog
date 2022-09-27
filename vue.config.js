const { defineConfig } = require("@vue/cli-service");
const IS_PROD = process.env.NODE_ENV === "production";
const path = require("path");
module.exports = defineConfig({
  publicPath: IS_PROD ? "" : "/",
  pages: {
    index: {
      entry: "src/main.ts",
      template: "public/index.html",
      // output as dist/index.html
      filename: "index.html",
      title: "zerotower的技术小屋",
    },
  },
  devServer: {
    host: "0.0.0.0",
    port: 9000,
    hot: true,
  },
  css: {
    loaderOptions: {
      css: {
        modules: {
          //[name] 是文件名 [local]是具体的类名
          //类名模块化，防止重复
          //auto这样设置就可以防止vue文件中的style标签定义的样式丢失了
          auto: /\.module\.scss$/i,
          // auto: true,
          // exportLocalsConvention: "asIs",
          // auto: true,
          localIdentName: IS_PROD ? "[hash:8]" : "[local]-[hash:base64:6]",
          // localIdentContext: path.resolve(__dirname, "src"),
          // namedExport: true,
        },
        // esModule: false,
        sourceMap: !IS_PROD,
      },
      sass: {
        // sassOptions: (loaderContext) => {
        //   const { resourcePath, rootContext } = loaderContext
        //   if (/\.vue/.test(resourcePath)) {
        //     console.log("resourcePath====>", resourcePath);
        //     console.log("loaderContext===>", loaderContext);
        //     process.abort()
        //   }
        // }
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
      if (IS_PROD) {
        config.output.filename("js/[name].[contenthash:8].js").end();
        config.output.chunkFilename("js/[name].[contenthash:8].js").end();
        config.plugin("extract-css").tap((args) => [
          {
            filename: `css/[name].[contenthash:8].css`,
            chunkFilename: `css/[name].[contenthash:8].css`,
          },
        ]);
      }
    }
  },
  productionSourceMap: false,
});
