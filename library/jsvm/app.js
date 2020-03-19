// app.js
var vm = new Vue({
    template: '<div>{{ msg }}</div>',
    serverPrefetch:function(){
        console.log(JSON.stringify(arguments));
    },
    computed: {
        msg: function () {
            return this.login ? "zhaopeng" : "guset"
        }
    },
    data: {
        login: false
    }
})
console.log('232323');
// 通过 `vue-server-renderer/basic.js` 暴露
renderVueComponentToString(vm, function (err, res) {
    console.log(res)
})