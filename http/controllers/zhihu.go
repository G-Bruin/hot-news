package controllers

import "net/http"

func Top(c *gin.Context) {
	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"

	list, _ := Curl("GET", url, "")
	//
	//test := list["data"].(map[string]interface{})
	//fmt.Println(test)

	//for _, value := range list {
	//	fmt.Println(value["fresh_text"])
	//}

	ReturnJson(c, http.StatusOK, "success", list)
}
