package rpc

import (
	"context"

	"github.com/UmaruCMS/auth-system/controller/user"
	pb "github.com/UmaruCMS/auth-system/rpc/base/auth"
)

type AuthServer struct{}

func (authServer *AuthServer) VerifyToken(ctx context.Context, rawToken *pb.Token) (*pb.Token, error) {
	t, err := user.ParseTokenString(rawToken.Raw)
	if err != nil {
		return nil, err
	}
	claims := t.Claims.(*user.CustomClaims)
	return &pb.Token{
		Raw:    rawToken.Raw,
		UserId: uint32(claims.ID),
		Valid:  t.Valid,
	}, nil
}
