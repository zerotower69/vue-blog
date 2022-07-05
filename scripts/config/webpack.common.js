const AntdDayjsWebpackPlugin = require("antd-dayjs-webpack-plugin");
const { CleanWebpackPlugin } = require("clean-webpack-plugin");
const CopyPlugin = require("copy-webpack-plugin");
const CompressPlugin = require("compression-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const webpack = require("webpack");
const WebpackBar = require("webpackbar");
const path = require("path");
// const Config = require("webpack-chain");
const {VueLoaderPlugin} =require('vue-loader')
const { ROOT_PATH } = require("../constant");
const { isDevelopment, isProduction } = require("../env");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
/**
 * resolve the dir src
 * @param {string} src
 * @returns string
 */
const resolve = (src) => {
  return path.resolve(ROOT_PATH, src);
};

/**
 * @type {(import 'webpack').Configuration}
 */
const commonConfig={
  //1.entry
  // context:resolve(ROOT_PATH,'/'),
  entry:{
    index:resolve('./src/main.ts'),
  },
  //2.output
  output:{
    // publicPath:"/static/",
    path:resolve('dist'),
    filename:"js/[name].[contenthash:8].js",
    // chunkFilename:'[name].[contenthash:8].js'
  },
  //3.modules
  module:{
    rules:[
      //deal vue files
      {
        test:/\.vue$/,
        use:'vue-loader',
        include:resolve('src')
      },
      //deal css files
      {
        test:/\.css$/,
        use:[
          "vue-style-loader",
          {
            loader:'css-loader',
            options:{
              sourceMap:isDevelopment,
              esModule:false,
            }
          }
        ]
      },
      //deal less files
      {
        test:/\.less$/,
        use:[
          'vue-style-loader',
          {
            loader:'css-loader',
            options:{
              sourceMap:isDevelopment,
            }
          },
          'less-loader'
      ]
      },
      //deal scss/sass files
      {
        test:/\.s(a|c)ss$/,
        use:[
          'style-loader',
          {
            loader:'css-loader',
            options:{
              sourceMap:isDevelopment
            }
          },
          'sass-loader'
        ]
      },
      //deal ts files
      {
        test:/\.tsx?$/,
        use:[
          {
            loader:'ts-loader',
            options:{
              appendTsSuffixTo: [/\.vue$/]
            }
          }
      ],
      exclude:resolve('node_modules')
      },
      //babel js files
      {
        test:/\.m?js$/,
        use:'babel-loader',
        exclude:resolve('node_modules')
      },
      //deal images files
      {
        test:/\.(bmp|gif|jpe?g|png)$/,
        use:[
          {
            loader:'url-loader',
            options:{
              limit: 8192,
              name: 'images/[name].[hash:7].[ext]',
              publicPath: './'
            }
          }
        ]
      },
      //deal fonts files
      {
        test:/\.(eot|svg|ttf|woff|woff2?)$/,
        use:'file-loader'
      }
    ]
  },
  //4.plugins
  plugins:[
    new VueLoaderPlugin(),
    new HtmlWebpackPlugin({
      template: resolve("./public/index.html"),
      filename: "index.html",
    //   chunks:["all"],
      inject: "body",
    }),
    new WebpackBar(),
    //copy plugin
    //clean plugin
    new CleanWebpackPlugin()
  ],
  //5.resolve
  resolve:{
    alias: {
      '@': path.resolve(ROOT_PATH, './src')
    },
    extensions:['.ts','.tsx','.vue','.js','.json'],
    fallback:{crypto:false}
  },
  //6.cache
  cache:{
    type:'filesystem',
    compression:'gzip'
  }
}
module.exports = commonConfig;
