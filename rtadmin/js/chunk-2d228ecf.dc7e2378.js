(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d228ecf"],{daea:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("Table",{attrs:{border:"",columns:t.columns,data:t.pages.data},scopedSlots:t._u([{key:"name",fn:function(e){var a=e.row;return[n("strong",[t._v(t._s(a.name))])]}},{key:"action",fn:function(e){var a=e.row,s=e.index;return[n("Button",{attrs:{type:"primary",size:"small",to:"/post/edit?id="+a.id}},[t._v(" 编辑 ")]),n("Poptip",{attrs:{confirm:"",placement:"left",title:"Are you sure you want to delete this item?"},on:{"on-ok":function(e){return t.deleteHandler(a,s)}}},[n("Button",{attrs:{size:"small"}},[t._v("Delete")])],1)]}}])}),n("br"),n("Page",{attrs:{total:t.pages.total},on:{"on-change":t.getData}})],1)},s=[],r=(n("a434"),n("d3b7"),n("96cf"),{data:function(){return{columns:[{title:"id",width:60,key:"id"},{title:"标题",key:"title"},{title:"发布时间",key:"created_at"},{title:"操作",slot:"action",width:150,align:"center"}],pages:{data:[],total:0}}},mounted:function(){this.getData(1)},methods:{getData:function(t){var e;return regeneratorRuntime.async((function(n){while(1)switch(n.prev=n.next){case 0:return e={page:t},n.next=3,regeneratorRuntime.awrap(this.$store.dispatch("posts/list",e));case 3:this.pages=n.sent;case 4:case"end":return n.stop()}}),null,this)},deleteHandler:function(t,e){this.$store.dispatch("posts/delete",t.id),this.pages.data.splice(e,1)}}}),i=r,o=n("2877"),l=Object(o["a"])(i,a,s,!1,null,null,null);e["default"]=l.exports}}]);
//# sourceMappingURL=chunk-2d228ecf.dc7e2378.js.map