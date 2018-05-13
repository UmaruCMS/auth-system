package rpc

import (
	"context"

	"github.com/UmaruCMS/auth-system/model"
	pb "github.com/UmaruCMS/auth-system/rpc/base/user"
)

type UserServer struct{}

func (userServer *UserServer) GetUserInfoByID(ctx context.Context, rawInfo *pb.UserInfo) (*pb.UserInfo, error) {
	userID := rawInfo.Id
	user := &model.User{}
	user, err := user.GetByID(uint(userID))
	if err != nil {
		return nil, err
	}
	return &pb.UserInfo{
		Id:    userID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
