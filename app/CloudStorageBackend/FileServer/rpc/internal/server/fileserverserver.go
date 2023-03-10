// Code generated by goctl. DO NOT EDIT!
// Source: fileServer.proto

package server

import (
	"context"

	"fileServer/rpc/internal/logic"
	"fileServer/rpc/internal/svc"
	"fileServer/rpc/pb"
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

// 就绪性探针
func (s *FileServerServer) Ping(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *FileServerServer) FindOne(ctx context.Context, in *pb.FindFileReq) (*pb.FileMetaInfo, error) {
	l := logic.NewFindOneLogic(ctx, s.svcCtx)
	return l.FindOne(in)
}

func (s *FileServerServer) QueryFiles(ctx context.Context, in *pb.QueryFileReq) (*pb.QueryFileRes, error) {
	l := logic.NewQueryFilesLogic(ctx, s.svcCtx)
	return l.QueryFiles(in)
}

func (s *FileServerServer) InertOne(ctx context.Context, in *pb.FileMetaInfo) (*pb.BaseRes, error) {
	l := logic.NewInertOneLogic(ctx, s.svcCtx)
	return l.InertOne(in)
}

func (s *FileServerServer) DeleteOne(ctx context.Context, in *pb.FileMetaInfo) (*pb.BaseRes, error) {
	l := logic.NewDeleteOneLogic(ctx, s.svcCtx)
	return l.DeleteOne(in)
}

func (s *FileServerServer) DeleteHard(ctx context.Context, in *pb.BaseTime) (*pb.BaseRes, error) {
	l := logic.NewDeleteHardLogic(ctx, s.svcCtx)
	return l.DeleteHard(in)
}
