package routers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/controllers"
	"github.com/zituocn/VMovie/controllers/admin"
)

func init() {

	//pages
	beego.Router("/", &controllers.IndexHandel{}, "*:Index")
	beego.Router("/m/:cid:int/", &controllers.IndexHandel{}, "*:List")
	beego.Router("/m/:cid:int/:page:int/", &controllers.IndexHandel{}, "*:List")
	beego.Router("/v/:id:int/", &controllers.IndexHandel{}, "*:Detail")
	beego.Router("/search/:key(.+)/", &controllers.IndexHandel{}, "*:Search")
	beego.Router("/search/:key(.+)/:page:int/", &controllers.IndexHandel{}, "*:Search")
	beego.Router("/json/", &controllers.IndexHandel{}, "*:Json")
	beego.Router("/p/:ename(.+)/", &controllers.IndexHandel{}, "*:Page")
	beego.Router("/today/", &controllers.IndexHandel{}, "*:Today")
	beego.Router("/22v.net.html", &controllers.IndexHandel{}, "*:Start")

	///error handel
	beego.ErrorController(&controllers.HttpErrorHandel{})

	//api json

	jns := beego.NewNamespace("/api",
		beego.NSRouter("/", &controllers.ApiHandel{}, "*:Index"),
		beego.NSRouter("/v/:id:int/", &controllers.ApiHandel{}, "*:Detail"),
		beego.NSRouter("/m/:cid:int/", &controllers.ApiHandel{}, "*:List"),
		beego.NSRouter("/m/:cid:int/:page:int/", &controllers.ApiHandel{}, "*:List"),
		beego.NSRouter("/search/:key(.+)/", &controllers.ApiHandel{}, "*:Search"),
		beego.NSRouter("/search/:key(.+)/:page:int/", &controllers.ApiHandel{}, "*:Search"),
		beego.NSRouter("/today/", &controllers.ApiHandel{}, "*:Today"),
	)
	beego.AddNamespace(jns)

	///admin
	admindir := beego.AppConfig.String("admindir")
	if len(admindir) == 0 {
		admindir = "admin"
	}
	ns := beego.NewNamespace(admindir,
		beego.NSRouter("/", &admin.LoginHandel{}, "*:Login"),
		beego.NSRouter("/logout", &admin.LoginHandel{}, "*:Logout"),
		beego.NSRouter("/main", &admin.IndexHandel{}, "*:Main"),
		beego.NSRouter("/left", &admin.IndexHandel{}, "*:Left"),
		beego.NSRouter("/right", &admin.IndexHandel{}, "*:Right"),

		//影片管理
		beego.NSRouter("movie/add", &admin.MovieHandel{}, "*:Add"),
		beego.NSRouter("movie/edit/:id:int/", &admin.MovieHandel{}, "*:Edit"),
		beego.NSRouter("movie/save", &admin.MovieHandel{}, "post:Save"),
		beego.NSRouter("movie/list", &admin.MovieHandel{}, "*:List"),
		beego.NSRouter("movie/list/:page:int/", &admin.MovieHandel{}, "*:List"),

		//下载管理
		beego.NSRouter("down/add/:mid:int/", &admin.DownaddrHandel{}, "*:Add"),
		beego.NSRouter("down/list", &admin.DownaddrHandel{}, "*:List"),
		beego.NSRouter("down/save/:ep:int/", &admin.DownaddrHandel{}, "*:Save"),

		//图片上传管理
		beego.NSRouter("upload/add", &admin.UploadHandel{}, "*:UpLoadPage"),
		beego.NSRouter("upload/qiniusave", &admin.UploadHandel{}, "*:QiniuUpLoadFile"),

		//用户
		beego.NSRouter("user/changepassword", &admin.UserHandel{}, "*:ChangePass"),
		beego.NSRouter("user/savepass", &admin.UserHandel{}, "*:SavePass"),

		//影片关系
		beego.NSRouter("relation/add", &admin.RelationHandel{}, "*:Add"),
		beego.NSRouter("relation/save", &admin.RelationHandel{}, "*:Save"),
		beego.NSRouter("relation/list", &admin.RelationHandel{}, "*:List"),
		beego.NSRouter("relation/list/:page:int/", &admin.RelationHandel{}, "*:List"),
		beego.NSRouter("relation/detail/:id:int/", &admin.RelationHandel{}, "*:Detail"),
		beego.NSRouter("relation/delete/:id:int/", &admin.RelationHandel{}, "*:Delete"),

		//专题管理
		beego.NSRouter("page/add", &admin.PageHandel{}, "*:Add"),
		beego.NSRouter("page/save", &admin.PageHandel{}, "*:Save"),
	)
	beego.AddNamespace(ns)
}
