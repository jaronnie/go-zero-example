package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) QueryPasswordByUsername(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Password: "123456",
	}, nil
}

func (l *GreetLogic) GetAllUser(req types.Request) (*types.AllUser, error) {
	// todo: add your logic here and delete this line

	return &types.AllUser {
		Alluser: []string{"go", "vue"},
	}, nil
}
