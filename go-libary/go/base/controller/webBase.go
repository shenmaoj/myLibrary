/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-24 09:52 
# @File : webBase.go
# @Description : 
*/
package controllers

import (
	"myLibrary/library/src/main/go/page"
	"strconv"
)

type WebBaseController struct {
	BaseController
}


func (this *WebBaseController) GetBasePageReq() (page.BasePageReq, error) {
	p := page.BasePageReq{}
	pageSizeStr := this.GetString("pageSize")
	pageNumStr := this.GetString("CurrentPage")
	p.PageSize = 10
	p.PageNum = 1
	if pageSizeStr != "" {
		pageSizeInt, e := strconv.Atoi(pageSizeStr)
		if nil != e {
			this.Log.Error("[QueryLastestBlocks] 请求参数pageSize错误,只可为整数")
			return p, e
		}
		p.PageSize = pageSizeInt
	}
	if pageNumStr != "" {
		pageNumInt, e := strconv.Atoi(pageNumStr)
		if nil != e {
			this.Log.Error("[QueryLastestBlocks] 请求参数pageNum错误,只可为整数")
			return p, e
		}
		p.PageNum = pageNumInt
	}

	return p,nil
}

func (this *WebBaseController)JSON(data interface{}){
}
