package util

import (
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"
	"errors"
	"reflect"
)

/**
maybe has some question
*/
func ConvertRpcFileMeta(info *pb.FileMetaInfo) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 10)
	refValueOf := reflect.ValueOf(*info)
	refTypeOf := reflect.TypeOf(*info)
	fieldNumber := refValueOf.NumField()
	for i := 0; i < fieldNumber; i++ {
		field := refValueOf.Field(i)
		fieldName := refTypeOf.Field(i).Name
		switch field.Kind() {
		case reflect.Invalid:
			err = errors.New("系统出错")
			break
		case reflect.Int, reflect.Int64:
			result[fieldName] = field.Int()
		case reflect.String:
			result[fieldName] = field.String()
		case reflect.Bool:
			result[fieldName] = field.Bool()
		case reflect.Float64, reflect.Float32:
			result[fieldName] = field.Float()
		default:
			continue
		}
	}
	return result, err
}

func ConvertRpcFileMetaList(info []*pb.FileMetaInfo) (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 10)
	for i := 0; i < len(info); i++ {
		result[i], err = ConvertRpcFileMeta(info[i])
		if err != nil {
			return nil, err
		}
	}
	return result, err
}
