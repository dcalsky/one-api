package render

import (
	"fmt"
	"github.com/bytedance/sonic"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/common"
)

func StringData(c *gin.Context, str string) {
	str = strings.TrimPrefix(str, "data: ")
	str = strings.TrimSuffix(str, "\r")
	c.Render(-1, common.CustomEvent{Data: "data: " + str})
	c.Writer.Flush()
}

func ObjectData(c *gin.Context, object interface{}) error {
	jsonData, err := sonic.MarshalString(object)
	if err != nil {
		return fmt.Errorf("error marshalling object: %w", err)
	}
	StringData(c, jsonData)
	return nil
}

func Done(c *gin.Context) {
	StringData(c, "[DONE]")
}
