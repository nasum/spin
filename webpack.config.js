const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path')

module.exports = {
  mode: "development",
  entry: path.resolve(__dirname, "frontend", "src", "index.ts"),
  output: {
    path: path.resolve(__dirname, "dist", "js")
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js"]
  },
  module: {
    rules:[
      {
        test: /\.tsx?$/, loader: "ts-loader"
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: "../index.html",
      template: path.resolve(__dirname, "frontend", "statics", "index.html")
    })
  ]
}
