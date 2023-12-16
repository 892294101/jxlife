module.exports = {
    lintOnSave: false,
    devServer: {
        port: 8080,
        //disableHostCheck: true
    },
    publicPath: '/',
    configureWebpack: {
        module: {
            rules: [{
                test: /\.mjs$/,
                include: /node_modules/,
                type: 'javascript/auto'
            }]
        }
    }
}
