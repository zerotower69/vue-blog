const webpack =require('webpack')
const {merge}=require('webpack-merge')
const CompressPlugin =require('compression-webpack-plugin')
const commonConfig =require('./webpack.common');

/**
 * @type {import('webpack').Configuration}
 */
const ProdConfig={
    mode:"production",
    optimization:{
        splitChunks:{  //split chunks to small
            chunks:'all'
        },
       runtimeChunk:'single'//单入口
    },
    plugins:[
        new CompressPlugin({
            algorithm:'gzip' //open gzip
        })
    ]
}

module.exports =merge(commonConfig,ProdConfig)

