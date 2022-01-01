package status

var ServerVersion string
var BuildTime string
var GitRef string

func VersionString() string {
	return ServerVersion
}

func InfoString() interface{} {
	resultMap := map[string]interface{}{"Version": ServerVersion, "BuildTime": BuildTime, "GitRef": GitRef}
	return resultMap
}
