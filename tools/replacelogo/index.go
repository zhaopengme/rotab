package replacelogo

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

func Mongo() {
	sqlTpl := `
db.create_template.insert(
{
  "adType": "PRODUCT_CATALOG_SALES",
  "channels": [
    "openx",
    "bidswitch",
    "yeahtargetergeniee",
    "pubmatic",
    "sonobi"
  ],
  "createDate": "2020-01-09",
  "createTime": 1578560759614,
  "creativeType": "googleHtml",
  "custom": {},
  "domainIds": [
    "${domainId}"
  ],
  "feed_num": 8,
  "feed_type": 1,
  "htmlContext": "<meta name=\"GCD\" content=\"YTk3ODQ3ZWZhN2I4NzZmMzBkNTEwYjJled4e0395df5aa56147dadb79f5c810f8\"/>\n  <meta charset=\"utf-8\">\n  <meta name=\"generator\" content=\"Google Web Designer 7.0.1.1101\">\n  <meta name=\"template\" content=\"Banner 3.0.0\">\n  <meta name=\"environment\" content=\"gwd-dv360\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <style>.gwd-inactive{visibility:hidden}gwd-page{display:block}</style>\n  <style>.gwd-pagedeck{position:relative;display:block}.gwd-pagedeck>.gwd-page.transparent{opacity:0}.gwd-pagedeck>.gwd-page{position:absolute;top:0;left:0;-webkit-transition-property:-webkit-transform,opacity;-moz-transition-property:transform,opacity;transition-property:transform,opacity}.gwd-pagedeck>.gwd-page.linear{transition-timing-function:linear}.gwd-pagedeck>.gwd-page.ease-in{transition-timing-function:ease-in}.gwd-pagedeck>.gwd-page.ease-out{transition-timing-function:ease-out}.gwd-pagedeck>.gwd-page.ease{transition-timing-function:ease}.gwd-pagedeck>.gwd-page.ease-in-out{transition-timing-function:ease-in-out}.linear *,.ease-in *,.ease-out *,.ease *,.ease-in-out *{-webkit-transform:translateZ(0);transform:translateZ(0)}</style>\n  <style>div[is=gwd-page].fs,gwd-page.fs{border:none}</style>\n  <style>gwd-image.scaled-proportionally>div.intermediate-element>img{background-repeat:no-repeat;background-position:center}gwd-image{display:inline-block}gwd-image>div.intermediate-element{width:100%;height:100%}gwd-image>div.intermediate-element>img{display:block;width:100%;height:100%}</style>\n  <style type=\"text/css\" id=\"gwd-lightbox-style\">.gwd-lightbox{overflow:hidden}</style>\n  <style type=\"text/css\" id=\"gwd-text-style\">p{margin:0px}h1{margin:0px}h2{margin:0px}h3{margin:0px}</style>\n  <style>.item-image{z-index:1;}.product-discount-img{z-index:2;}.product-discount{z-index:3;}</style>\n  <link rel=\"stylesheet\" href=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/css.css\">\n\n        <!-- base -->\n        <script data-exports-type=\"gwd-google-ad\" src=\"https://d1b00yffurhml5.cloudfront.net/tpl/js/Enabler.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/google.min.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/gwd_webcomponents_min.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/gwdpage_min.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/gwdpagedeck_min.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/gwdgooglead_min.js\"></script>\n        <script src=\"https://d1b00yffurhml5.cloudfront.net/tpl/google-web/google-base/gwdimage_min.js\"></script>\n  <gwd-google-ad id=\"gwd-ad\" polite-load=\"\">\n    <gwd-metric-configuration></gwd-metric-configuration>\n    <gwd-pagedeck class=\"gwd-page-container\" id=\"pagedeck\">\n      <gwd-page id=\"page1\" class=\"gwd-page-wrapper gwd-page-size gwd-lightbox\" data-gwd-width=\"100%\" data-gwd-height=\"100%\">\n        <div class=\"gwd-page-content gwd-page-size\">\n          <div class=\"gwd-div-1ln5\" id=\"gwd-feed-price-wrapper\" red-border-carousel=\"red\">\n            <gwd-image id=\"001\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/001.png\" scaling=\"stretch\" class=\"item-image gwd-image-kige\"></gwd-image>\n            <gwd-image id=\"002\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/002.png\" scaling=\"stretch\" class=\"item-image gwd-image-86ew\"></gwd-image>\n            <gwd-image id=\"003\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/003.png\" scaling=\"stretch\" class=\"item-image gwd-image-1qcp\"></gwd-image>\n            <gwd-image id=\"004\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/004.png\" scaling=\"stretch\" class=\"item-image gwd-image-1d33\"></gwd-image>\n            <gwd-image id=\"005\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/005.png\" scaling=\"stretch\" class=\"item-image gwd-image-15w1\"></gwd-image>\n            <gwd-image id=\"006\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/006.png\" scaling=\"stretch\" class=\"item-image gwd-image-5m7r\"></gwd-image>\n            <gwd-image id=\"007\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/007.png\" scaling=\"stretch\" class=\"item-image gwd-image-2whd\"></gwd-image>\n            <gwd-image id=\"008\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/008.png\" scaling=\"stretch\" class=\"item-image gwd-image-1rsv\"></gwd-image>\n            <gwd-image id=\"009\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/009.png\" scaling=\"stretch\" class=\"item-image gwd-image-nxph\"></gwd-image>\n            <gwd-image id=\"010\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/010.png\" scaling=\"stretch\" class=\"item-image gwd-image-1hxs\"></gwd-image>\n            <gwd-image id=\"011\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/011.png\" scaling=\"stretch\" class=\"item-image gwd-image-o2kb\"></gwd-image>\n            <gwd-image id=\"012\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/012.png\" scaling=\"stretch\" class=\"item-image gwd-image-scgq\"></gwd-image>\n            <gwd-image id=\"013\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/013.png\" scaling=\"stretch\" class=\"item-image gwd-image-1lh3\"></gwd-image>\n          </div>\n          <gwd-image id=\"patpat-\" source=\"${domainlogo}\" scaling=\"stretch\" class=\"gwd-image-uwz3\"></gwd-image>\n          <div class=\"gwd-div-1595\" id=\"gwd-price\">\n            <span class=\"product-price gwd-span-nbw0 gwd-span-kjfp\" id=\"$1999-01\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-1sfz\" id=\"$1999-02\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-18ir\" id=\"$1999-03\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-qgx1 gwd-span-huls\" id=\"$1999-04\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-qgx1 gwd-span-2gc0 gwd-span-l732\" id=\"$1999-05\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-c8j3\" id=\"$1999-06\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-1ssf\" id=\"$1999-07\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-lvja\" id=\"$1999-08\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-qgx1 gwd-span-10tk\" id=\"$1999-09\">$19.99</span>\n            <span class=\"product-price gwd-span-nbw0 gwd-span-m502 gwd-span-jjm5 gwd-span-qgx1 gwd-span-2gc0 gwd-span-eyop\" id=\"$1999-10\">$19.99</span>\n          </div>\n          <div class=\"gwd-div-pi9n\">\n            <gwd-image id=\"label_1\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-zux5\"></gwd-image>\n            <gwd-image id=\"label_2\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-1umf\"></gwd-image>\n            <gwd-image id=\"label_3\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-1bg3\"></gwd-image>\n            <gwd-image id=\"label_4\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-3l6c\"></gwd-image>\n            <gwd-image id=\"label_5\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-1azu gwd-image-1eyo\"></gwd-image>\n            <gwd-image id=\"label_6\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-1azu gwd-image-16wm gwd-image-89ra\"></gwd-image>\n            <gwd-image id=\"label_7\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-1azu gwd-image-16wm gwd-image-1xu5 gwd-image-1yj5\"></gwd-image>\n            <gwd-image id=\"label_8\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-1azu gwd-image-16wm gwd-image-1xu5 gwd-image-9o7f gwd-image-vjp0\"></gwd-image>\n            <gwd-image id=\"label_9\" source=\"https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/____2-02.png\" scaling=\"stretch\" class=\"product-discount-img gwd-image-1d6u gwd-image-ks4j gwd-image-jpvx gwd-image-1bkz gwd-image-19ay gwd-image-tx4f gwd-image-1azu gwd-image-16wm gwd-image-1xu5 gwd-image-9o7f gwd-image-y7c4\"></gwd-image>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1ky4\" id=\"-20_1\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-g769\" id=\"-20_2\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-1v2i gwd-span-no1v\" id=\"-20_3\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-vnzj gwd-span-ul0r\" id=\"-20_4\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-dzc7 gwd-span-vnst\" id=\"-20_5\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-1v2i gwd-span-1yns gwd-span-1yyg\" id=\"-20_6\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-vnzj gwd-span-2ozm\" id=\"-20_7\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-dzc7 gwd-span-11oi\" id=\"-20_8\">-20%</span>\n            <span class=\"product-discount gwd-span-iqep gwd-span-1qts gwd-span-1v2i gwd-span-1yns gwd-span-1th5\" id=\"-20_9\">-20%</span>\n          </div>\n        </div>\n      </gwd-page>\n    </gwd-pagedeck>\n  </gwd-google-ad>\n  <script type=\"text/javascript\" id=\"gwd-init-code\">\n    (function() {\n      var gwdAd = document.getElementById('gwd-ad');\n\n      /**\n       * Handles the DOMContentLoaded event. The DOMContentLoaded event is\n       * fired when the document has been completely loaded and parsed.\n       */\n      function handleDomContentLoaded(event) {\n\n      }\n\n      /**\n       * Handles the WebComponentsReady event. This event is fired when all\n       * custom elements have been registered and upgraded.\n       */\n      function handleWebComponentsReady(event) {\n        // Start the Ad lifecycle.\n        setTimeout(function() {\n          gwdAd.initAd();\n        }, 0);\n      }\n\n      /**\n       * Handles the event that is dispatched after the Ad has been\n       * initialized and before the default page of the Ad is shown.\n       */\n      function handleAdInitialized(event) {}\n\n      window.addEventListener('DOMContentLoaded',\n        handleDomContentLoaded, false);\n      window.addEventListener('WebComponentsReady',\n        handleWebComponentsReady, false);\n      window.addEventListener('adinitialized',\n        handleAdInitialized, false);\n    })();\n  </script>\n\n\n<script data-exports-type=\"gwd-studio-registration\">function StudioExports() {\n}</script><script type=\"text/gwd-admetadata\">{\"version\":1,\"type\":\"GoogleAd\",\"format\":\"\",\"template\":\"Banner 3.0.0\",\"politeload\":true,\"fullscreen\":false,\"counters\":[],\"timers\":[],\"exits\":[],\"creativeProperties\":{\"minWidth\":0,\"minHeight\":0,\"maxWidth\":0,\"maxHeight\":0},\"components\":[\"gwd-google-ad\",\"gwd-image\",\"gwd-page\",\"gwd-pagedeck\"],\"responsive\":true}</script>\n\n\n<script type=\"text/javascript\">\n  var carouselImgNum = {\n    '728-90': 5,\n    '250-250': 2,\n    '300-250': 2,\n    '970-250': 4,\n    '336-280': 2,\n    '120-600': 3,\n    '160-600': 3,\n    '300-600': 2,\n  }\n\n  var carouselBorderColor = \"red\";\n</script>",
  "id": "2006-${id}",
  "name": "2020元旦模板-${domainName}-非google",
  "preview_url": "https://d1b00yffurhml5.cloudfront.net/tpl/20191226/01/preview.png",
  "score": 90,
  "sizeList": [
    "728_90",
    "250_250",
    "300_250",
    "970_250",
    "336_280",
    "120_600",
    "300_600"
  ],
  "tags": "非 google",
  "templateTypes": [
    "holiday"
  ],
  "types": [
    "clothes"
  ],
  "updateTime": "2020-01-09T09:05:59.621Z",
  "delete": 0
}
)
`

	dataJson := `
[
{
  "_id": "5962ed388efcf444900384fb",
  "domain": "lovelywholesale.com",
  "id": "b95a9d841679483da2b19f9d1202e38c",
  "logo":"lovelywhole-.png"
},
{
  "_id": "5e056fa7d73a9009672b0540",
  "id": "7cb5f7e9b2e247f39e4c5db8e2ae696a",
  "domain": "geekbuying.com",
  "logo":"geekbuy.png"
},
{
  "_id": "5d6f696c0f25215c7cd6d025",
  "id": "5184416cc53348e787f3395e20b362f6",
  "domain": "africanmall.com",
  "logo":"Afri.png"
},
{
  "_id": "5d5a46d90f25215c7c102229",
  "id": "ac975ae434c04024b1a66ba30e82374e",
  "domain": "bellewholesale.com",
  "logo":"bellewholesale-.png"
},
{
  "_id": "5cff69470f25215c7c248df1",
  "id": "ebc0bc9a334f4eddab9ecb8ee93b4ce9",
  "domain": "bc-style.jp",
  "logo":"bc-style.png"
},
{
  "_id": "5cc7b0bd0f25215c7cf0a42b",
  "id": "7ff912b236544416a68213a6fdfc6f44",
  "domain": "noracora.com",
  "logo":"NORACORA.png"
},
{
  "_id": "5bd2c238e6de2eea5c9b66b4",
  "id": "5136f65b25e54309bc77c56314c53668",
  "domain": "tijneyewear.com",
  "logo":"tijn-.png"
},
{
  "_id": "5bd27823e6de2eea5cbd54c8",
  "id": "ec8b37a91a634e5288ad6a9896fd95a0",
  "domain": "patpat.com",
  "logo":"patpat-.png"
},
{
  "_id": "5bc3fc3de6de2eea5cf6ec75",
  "id": "9484ace7a8464765abdb1a4f143717ae",
  "domain": "boutiquefeel.com",
  "logo":"bq-.png"
},
{
  "_id": "5bc3fbf9e6de2eea5cf32d46",
  "id": "5ef3f218fcd241d18c7d6530cde4a582",
  "domain": "ivrose.com",
  "logo":"IVRSE-.png"
},
{
  "_id": "5b92432ce6de2eea5c490584",
  "id": "f30c69753176479aaeb05a9347a2fceb",
  "domain": "zivame.com",
  "logo":"zivame.png"
},
{
  "_id": "5b6ba36fe6de2eea5c25a2bc",
  "id": "06eeef1ad96f45c69b6edba5e4c66b33",
  "domain": "chicme.com",
  "logo":"chicme-.png"
},
{
  "_id": "5a37838953cd2907fa27ea0e",
  "domain": "stylewe.com",
  "id": "e916cef2d0554505b7fe6ddd0add8154",
  "logo":"stylewe-.png"
},
{
  "_id": "5991a1539429480b342a185b",
  "domain": "www.soufeel.com",
  "id": "5021cc7bab4d480a872f61f526e33754",
  "logo":"soufeel-.png"
}
]
`
	domains, _ := gjson.LoadContent(dataJson)
	for _, domain := range domains.ToArray() {
		println("========================================================")
		domainId := gconv.String(domain.(map[string]interface{})["id"])
		domainName := gconv.String(domain.(map[string]interface{})["domain"])
		domainlogo := "https://d1b00yffurhml5.cloudfront.net/tpl/google-images/rtb-logo-1/" + gconv.String(domain.(map[string]interface{})["logo"])
		sql := gstr.Replace(sqlTpl, "${domainId}", domainId)
		sql = gstr.Replace(sql, "${domainName}", domainName)
		sql = gstr.Replace(sql, "${id}", gconv.String(grand.Intn(99999)))
		sql = gstr.Replace(sql, "${domainlogo}", domainlogo)
		gfile.PutContents("/Users/zhaopeng/Desktop/"+domainName+".txt", sql)
	}

}
