const path = require("path");
function resolve(dir) {
  return path.join(__dirname, dir);
}

module.exports = {
  chainWebpack: config => {
    config.resolve.alias
      .set("assets", resolve("src/assets"))
      .set("variables", resolve("src/styles/variables.styl"));
    config.module
      .rule("pdf")
      .test(/\.pdf/)
      .use("")
      .loader("file-loader");
  }
};
