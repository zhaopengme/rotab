package upload

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

func TestUpload(t *testing.T) {
	file := "/Users/zhaopeng/Downloads/geekbuy.png"
	params := g.Map{
		"reqtype":      "fileupload",
		"fileToUpload": "@file:" + file,
	}
	result := ghttp.PostContent("https://catbox.moe/user/api.php", params)
	println(result)
}
