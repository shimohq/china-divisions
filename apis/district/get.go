package district

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

const districtDataDir = "./district_data/dist/"

type provinceModel struct {
	Code     string      `json:"code"`
	Name     string      `json:"name"`
	Children []cityModel `json:"children"`
}

type cityModel struct {
	Code     string      `json:"code"`
	Name     string      `json:"name"`
	Children []areaModel `json:"children"`
}

type areaModel struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Get 行政区划 Http Get 方法
// @Summary 获得行政区信息
// @Description 最多能获取到区级别的联动信息
// @Tags district
// @Accept  json
// @Produce  json
// @summary 如果不传入 code 或 code 无效，会根据 subdistrict 来返回所有的行政区数据；
// 传入 code 情况下，会根据 code 的长度和 subdistrict 来返回行政区数据，例如 code 340102 subdistrict 2 仅返回瑶海区
// @Param code query string false "行政区代码"
// @Param subdistrict query integer false "string default" default(1)
// @Success 200
// @Router /district [get]
func Get(c *gin.Context) {
	fileSubdistrict := map[string]string{
		"1": "provinces.json",
		"2": "pc-code.json",
		"3": "pca-code.json",
	}
	var message interface{}
	code := c.Query("code")
	subdistrict := c.Query("subdistrict")
	if subdistrict == "" {
		subdistrict = "1"
	}
	subdistrictNum, err := strconv.Atoi(subdistrict)
	if err != nil {
		c.JSON(400, gin.H{
			"errmessage": "subdistrict must be number",
		})
	}
	if len(code) >= subdistrictNum*2 {
		code = string([]byte(code)[:subdistrictNum*2])
	}

	provinces := []provinceModel{}
	err = readFile(districtDataDir+fileSubdistrict[subdistrict], &provinces)
	if err != nil {
		log.Panicln(err)
	}

	message = provinces
	provinceCode := string([]byte(code)[:2])
	for _, province := range provinces {
		if provinceCode == province.Code {
			if len(code) == 2 {
				message = province
				break
			}

			if len(code) >= 4 {
				cityCode := string([]byte(code)[0:4])
				for _, city := range province.Children {
					if cityCode == city.Code {
						if len(code) == 4 {
							message = city
							break
						}

						if len(code) == 6 {
							areaCode := string([]byte(code)[0:6])
							for _, area := range city.Children {
								if areaCode == area.Code {
									message = area
									break
								}
							}
						}
					}
				}
			}
		}
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

func readFile(fileName string, data interface{}) error {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
