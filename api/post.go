package api

import (
	"errors"
	"go/common"
	"go/models"
	"go/service"
	"go/utils"
	"net/http"
	"strconv"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id，判断是否登录了
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已经过期"))
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		CategoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(int)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			CategoryId,
			uid,
			-1,
			postType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
	}

}
