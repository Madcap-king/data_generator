package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"mongo_data/constants"
	"mongo_data/model"
	"mongo_data/services"
	"net/http"
)

func DataCreateHandle(c *gin.Context) {
	data, _ := c.GetRawData()
	var reqVO model.RequestVO
	err := json.Unmarshal(data, &reqVO)
	if err != nil || !isValid(&reqVO) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	err = services.DataCreate(&reqVO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func isValid(m *model.RequestVO) bool {
	if len(m.DB) <= 0 || len(m.Col) <= 0 || len(m.Content) <= 0 || m.Rows <= 0 {
		return false
	}

	if !checkParserType(m) || !checkGeneratorType(m) {
		return false
	}

	// 校验content是否符合json规范
	var tmpMap map[string]interface{}
	err := json.Unmarshal([]byte(m.Content), &tmpMap)
	if err != nil {
		return false
	}
	return true
}

func checkGeneratorType(m *model.RequestVO) bool {
	if m.GenerateType != constants.Random && m.GenerateType != constants.WordLib {
		return false
	}

	return true
}

func checkParserType(m *model.RequestVO) bool {
	if m.Type != constants.JSON && m.Type != constants.XML {
		return false
	}

	return true
}
