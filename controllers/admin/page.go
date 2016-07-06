package admin

import (
	"github.com/zituocn/VMovie/models"
	//"strconv"
	"strings"
)

type PageHandel struct {
	baseController
}

func (this *PageHandel) Add() {
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/pageadd.html"
}

//保存专题
func (this *PageHandel) Save() {
	var (
		id          int64
		name        string
		ename       string
		content     string
		status      int64
		title       string
		keywords    string
		description string

		info models.PageInfo
		err  error
	)

	name = strings.TrimSpace(this.GetString("name"))
	ename = strings.TrimSpace(this.GetString("ename"))
	content = strings.TrimSpace(this.GetString("content"))
	title = strings.TrimSpace(this.GetString("title"))
	keywords = strings.TrimSpace(this.GetString("keywords"))
	description = strings.TrimSpace(this.GetString("description"))
	status, _ = this.GetInt64("status")
	id, _ = this.GetInt64("id")
	if len(name) == 0 || len(ename) == 0 || len(content) == 0 || len(title) == 0 || len(description) == 0 {
		this.showmsg("对不起,带*号的项必须填写...")
	}

	info.Name = name
	info.Ename = ename
	info.Editor = this.nickname
	info.Title = title
	info.Content = content
	info.Keywords = keywords
	info.Description = description
	info.Status = status
	if id > 0 {
		info.Id = id
		err = info.Update("name", "content", "title", "keywords", "description", "status")
	} else {
		err = info.Insert()
	}
	if err != nil {
		this.showmsg("保存出错，错误信息：" + err.Error())
	} else {
		this.showmsg("专题数据保存成功...", this.admindir+"page/add")
	}
}
