package pb

import (
	"errors"
	"fileServer/model/PojoDB/fileMetaPojo"
	"fmt"
	"time"
)

func ConvertFileMetaInfo(filemeta *FileMetaInfo) (result *fileMetaPojo.Filemetatable, err error) {
	var createTime, updateTime, deleteTime time.Time
	createStr := filemeta.CreateTime
	updateStr := filemeta.UpdateTime
	deleteStr := filemeta.DeleteTime
	if createStr != "" {
		createTime, err = time.Parse(time.RFC850, createStr)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("convert error:%s", createStr))
		}
	}
	if updateStr != "" {
		updateTime, err = time.Parse(time.RFC850, updateStr)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("convert error:%s", updateStr))
		}
	}
	if deleteStr != "" {
		deleteTime, err = time.Parse(time.RFC850, deleteStr)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("convert error:%s", deleteStr))
		}
	}
	isDirInt := 0
	if filemeta.IsDir {
		isDirInt = 1
	}
	deleteTime, err = time.Parse(time.RFC850, filemeta.DeleteTime)
	result = &fileMetaPojo.Filemetatable{
		Creator:     filemeta.Creator,
		CreateGroup: filemeta.CreateGroup,
		Name:        filemeta.Name,
		CreateTime:  createTime,
		Authority:   filemeta.Authority,
		TypeOf:      filemeta.TypeOf,
		UpdateTime:  updateTime,
		Size:        filemeta.Size,
		IsDir:       isDirInt,
		Description: filemeta.Description,
	}
	result.DeleteTime.Time = deleteTime
	return result, err
}
func GetFileMetaInfo(filemetatable *fileMetaPojo.Filemetatable) (result *FileMetaInfo) {
	result = &FileMetaInfo{

		Creator:     filemetatable.Creator,
		CreateGroup: filemetatable.CreateGroup,
		Name:        filemetatable.Name,
		CreateTime:  filemetatable.CreateTime.String(),
		Authority:   filemetatable.Authority,
		TypeOf:      filemetatable.TypeOf,
		UpdateTime:  filemetatable.UpdateTime.String(),
		Size:        filemetatable.Size,
		Description: filemetatable.Description,
	}
	if filemetatable.IsDir == 0 {
		result.IsDir = true
	} else {
		result.IsDir = false
	}
	deleteTime := filemetatable.DeleteTime
	if deleteTime.Valid {
		result.DeleteTime = deleteTime.Time.String()
	}
	return result
}
func GetFileMetaList(filemetatable []*fileMetaPojo.Filemetatable) (result []*FileMetaInfo) {
	result = make([]*FileMetaInfo, len(filemetatable))
	for index, value := range filemetatable {
		result[index] = GetFileMetaInfo(value)
	}
	return result
}

//原型模式
func NewDefaultBaseRes() *BaseRes {
	return &BaseRes{Result: "true", Message: "处理成功"}
}

func (br *BaseRes) GetFailedRes(message string) {
	br.Result = "false"
	br.Message = message
}
