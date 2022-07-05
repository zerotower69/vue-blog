module.exports =
{
    "presets": [
      [
        "@babel/preset-env",
        {
          "modules": false
        }
      ],
      "@babel/preset-typescript"
    ],
    "plugins": [
      [
        "@babel/plugin-transform-runtime",
        {
          "corejs": {
            "version": 3,
            "proposals": true
          }
        }
      ],
      [
        "import",
        {
          "libraryName": "antd-design-vue",
          "libraryDirectory": "es",
          "style": true // `style: true` 会加载 less 文件
        }
      ],
      ["@babel/plugin-syntax-dynamic-import"]
    ]
  }
  