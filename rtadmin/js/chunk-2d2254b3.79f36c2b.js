(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d2254b3"],{e47e:function(r,s,e){"use strict";e.r(s);var o=function(){var r=this,s=r.$createElement,e=r._self._c||s;return e("div",[e("h3",[r._v("修改密码")]),e("br"),e("Form",{ref:"form",attrs:{model:r.form,rules:r.rules,"label-width":120}},[e("FormItem",{attrs:{label:"原密码",prop:"originPassword"}},[e("i-input",{attrs:{type:"password"},model:{value:r.form.originPassword,callback:function(s){r.$set(r.form,"originPassword",s)},expression:"form.originPassword"}})],1),e("FormItem",{attrs:{label:"新密码",prop:"password"}},[e("i-input",{attrs:{type:"password"},model:{value:r.form.password,callback:function(s){r.$set(r.form,"password",s)},expression:"form.password"}})],1),e("FormItem",{attrs:{label:"重复密码",prop:"password2"}},[e("i-input",{attrs:{type:"password"},model:{value:r.form.password2,callback:function(s){r.$set(r.form,"password2",s)},expression:"form.password2"}})],1),e("FormItem",[e("Button",{attrs:{type:"primary"},on:{click:r.submitHandler}},[r._v("更新")])],1)],1)],1)},t=[],a=(e("d3b7"),e("96cf"),{data:function(){var r=this,s=function(s,e,o){""===e?o(new Error("Please enter your password again")):r.form.password!==r.form.password2?o(new Error("The two input passwords do not match!")):o()};return{form:{originPassword:"",password:"",password2:""},rules:{originPassword:[{required:!0,message:"The password cannot be empty",trigger:"blur"}],password:[{required:!0,message:"The password cannot be empty",trigger:"blur"}],password2:[{required:!0,message:"The password cannot be empty",trigger:"blur"},{validator:s,trigger:"blur"}]}}},methods:{submitHandler:function(){var r=this;this.$refs.form.validate((function(s){var e;return regeneratorRuntime.async((function(o){while(1)switch(o.prev=o.next){case 0:if(!s){o.next=5;break}return o.next=3,regeneratorRuntime.awrap(r.$store.dispatch("user/updatePassword",r.form));case 3:e=o.sent,r.$kit.isBlank(e)&&(r.$Message.success("更新成功!"),r.$refs.form.resetFields());case 5:case"end":return o.stop()}}))}))}}}),n=a,i=e("2877"),d=Object(i["a"])(n,o,t,!1,null,null,null);s["default"]=d.exports}}]);
//# sourceMappingURL=chunk-2d2254b3.79f36c2b.js.map