package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Clear(c *gin.Context) {
	fileBytes, err := os.ReadFile(BasePath + "maplist.txt")
	if err != nil {
		c.String(http.StatusInternalServerError, "获取地图记录文件失败")
		return
	}

	fileList := strings.Split(string(fileBytes), "\n")
	fileList = fileList[:len(fileList)-1]
	errFileList := []string{}
	for _, file := range fileList {
		if err := os.Remove(BasePath + file); err != nil {
			errFileList = append(errFileList, file)
		}
	}

	if len(errFileList) > 0 {
		os.WriteFile(BasePath+"maplist.txt", []byte(strings.Join(errFileList, "\n")), 0666)
		c.String(http.StatusInternalServerError, "以下文件删除失败："+strings.Join(errFileList, ","))
		return
	}

	if err := os.Remove(BasePath + "maplist.txt"); err != nil {
		c.String(http.StatusInternalServerError, "删除地图记录文件失败")
		return
	}

	c.String(http.StatusOK, "清空成功！")
}
