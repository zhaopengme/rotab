package controller

import (
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"io"
)

type FileCtr struct{}

func init() {
	g.Server().BindObject("/file", new(FileCtr))
}

func (o *FileCtr) Upload(r *ghttp.Request) {
	f, h, e := r.FormFile("file")
	if e != nil {
		response.JsonFailError(r, e, nil)
	}
	defer f.Close()
	filename := gfile.Basename(h.Filename)
	savePath := gfile.Join(gfile.TempDir(), filename)
	file, err := gfile.Create(savePath)
	if err != nil {
		response.JsonFailError(r, e, nil)
	}
	defer file.Close()
	if _, err := io.Copy(file, f); err != nil {
		response.JsonFailError(r, e, nil)
	}

	params := g.Map{
		"reqtype":      "fileupload",
		"fileToUpload": "@file:" + savePath,
	}

	result := ghttp.PostContent("https://catbox.moe/user/api.php", params)
	gfile.Remove(savePath)
	response.JsonOK(r, "", g.Map{"name": filename, "url": result})
}
