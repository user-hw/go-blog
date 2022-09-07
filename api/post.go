package api

import (
	"errors"
	"fmt"
	"go/common"
	"go/dao"
	"go/models"
	"go/service"
	"go/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求"))
		return
	}
	post, err := dao.GetPostByPid(pid)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	common.Success(w, post)

}

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
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			CategoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		fmt.Println(post.Markdown)
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		//update
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		CategoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		pidFloat := params["pid"].(float64)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			CategoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		fmt.Println(post.Markdown)
		service.UpdatePost(post)
		common.Success(w, post)
	}

}
