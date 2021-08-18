package form

type VpsInit struct {
	Name string `json:"name" form:"name" binding:"required"`
	Port uint   `json:"port" form:"port" binding:"required"`
}

type HeartBeat struct {
	Type     string
	TargetIp string
	Time     int64
}
