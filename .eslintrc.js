module.exports = {
  root: true,

  env: {
    node: true,
  },

  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/typescript/recommended",
    "plugin:prettier/recommended",
  ],

  parserOptions: {
    ecmaVersion: 2020,
  },

  rules: {
    "no-console": "off",
    "no-debugger": "off",
    "vue/no-unused-components": "warn",
    "vue/no-unused-properties": "warn",
    "vue/no-unused-refs": "warn",
    "vue/no-use-computed-property-like-method": "warn",
    "vue/multi-word-component-names": "off",
    "@typescript-eslint/no-empty-function": "warn",
    "@typescript-eslint/ban-ts-comment": "off",
    "@typescript-eslint/no-var-requires": "off",
    "linebreak-style": [0, "off", "windows"],
  },

  overrides: [
    {
      files: [
        "**/__tests__/*.{j,t}s?(x)",
        "**/tests/unit/**/*.spec.{j,t}s?(x)",
      ],
      env: {
        jest: true,
      },
    },
  ],
};
