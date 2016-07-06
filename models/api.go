package models

///详情页输出结构
type ApiDetailInfo struct {
	Minfo    *MovieInfo      //影片信息
	SameList []*MovieInfo    //类似影片
	DownList []*DownAddrInfo //下载地址列表
	Cinfo    *MovieClassInfo //分类信息
}

//今日更新列表
type ApiTodayInfo struct {
	List []*MovieInfo //影片列表
}

//列表页输出模型
type ApiListInfo struct {
	Cinfo       *MovieClassInfo //分类信息
	MList       []*MovieInfo    //影片列表
	Page        int64           //当前页码
	RecordCount int64           //记当总数
}

type ApiIndexMovieList struct {
	MList []*MovieInfo //影片各种小列表
}

//首页输出模型
type ApiIndexInfo struct {
	List []*ApiIndexMovieList
}

//搜索模型
type ApiSearchInfo struct {
	List        []*MovieInfo //影片列表
	Page        int64
	RecordCount int64
}
