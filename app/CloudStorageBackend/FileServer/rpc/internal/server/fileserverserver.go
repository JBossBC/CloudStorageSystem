// Code generated by goctl. DO NOT EDIT!
// Source: fileServer.proto

package server

import (
	"context"

	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/logic"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/internal/svc"
	"cloudStorageSystem/app/CloudStorageBackend/FileServer/rpc/pb"
)

type FileServerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFileServerServer
}

func NewFileServerServer(svcCtx *svc.ServiceContext) *FileServerServer {
	return &FileServerServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServerServer) FindOne(ctx context.Context, in *pb.FindFileReq) (*pb.FileMetaInfo, error) {
	l := logic.NewFindOneLogic(ctx, s.svcCtx)
	return l.FindOne(in)
}

func (s *FileServerServer) QueryFiles(ctx context.Context, in *pb.FindFileReq) (*pb.QueryFileRes, error) {
	l := logic.NewQueryFilesLogic(ctx, s.svcCtx)
	return l.QueryFiles(in)
}
