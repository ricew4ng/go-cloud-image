package aliyun

type Region struct {
	Label    string // 地域名称
	EndPoint string // 端点
	RegionID string // 地域ID
}

type Image struct {
	Label   string // 镜像名称
	ImageID string // 镜像ID
	OS      string // 系统名称
	/*
		Creating：镜像正在创建中。
		Waiting：多任务排队中。
		Available（默认）：您可以使用的镜像。
		UnAvailable：您不能使用的镜像。
		CreateFailed：创建失败的镜像。
		Deprecated：已弃用的镜像。
	*/
	Status string // 状态
}
