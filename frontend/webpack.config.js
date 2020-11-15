const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const path = require('path');

module.exports = {
  mode: process.env.NODE_ENV,
  entry: {
    app: ['./src/index.js', './src/style.scss'],
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, '../assets/dist'),
  },
  module: {
    rules: [
      {
        test: /\.scss$/,
        use: [
          process.env.NODE_ENV !== 'production' ? 'style-loader' : MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader',
        ],
      },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].css',
    }),
  ],
};