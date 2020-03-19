package view

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestMergeObj(t *testing.T) {
	MergeObj("post", g.Map{"id": "id"}, "comments", g.Map{"comments": "comments"})
}
