package application

import "turato.com/bdntoy/service"

//Sessions 会话列表
func Sessions() ([]*service.BdHistorySessionRecord, error) {
	return getService().GetHistorySession()
}

// ShareGroups 通讯录-群组列表
func ShareGroups() ([]*service.ShareGroupInfo, error) {
	return getService().GetShareGroups()
}

// FileLibraries 文件库列表
func FileLibraries(gid string) (*service.GetShareGroupFileListRsp, error) {
	return getService().GetShareGroupFileList(gid)
}

// AllFiles 获取文件库的所有文件列表
func AllFiles(param service.ShareInfoParam) ([]*service.ShareRecord, error) {
	return getService().GetAllShareFiles(param)
}

// GetShareFiles 获取信息1层
func GetShareFiles(param service.ShareInfoParam) ([]*service.ShareRecord, error) {
	r, err := getService().GetShareFiles(param)
	if err != nil {
		return nil, err
	}
	return r.Records, nil
}
