// Code generated by hertz generator.

package api

import (
	"context"

	api "go-web/api/biz/model/api"
	"go-web/api/biz/rpc"
	"go-web/api/kitex_gen/demouser"

	"github.com/cloudwego/hertz/pkg/app"
	pconsts "github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateUser .
// @router /v2/user/register [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(pconsts.StatusBadRequest, err.Error())
		return
	}

	err = rpc.CreateUser(context.Background(), &demouser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.String(pconsts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateUserResponse)

	resp.BaseResp = new(api.BaseResp)
	resp.BaseResp.Code = 0
	resp.BaseResp.Message = "success"

	c.JSON(pconsts.StatusOK, resp)
}

// CheckUser .
// @router /v2/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CheckUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(pconsts.StatusBadRequest, err.Error())
		return
	}

	// TODO: check user, use Jwt middleware
	_, err = rpc.CheckUser(context.Background(), &demouser.CheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.String(pconsts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CreateUserResponse)

	resp.BaseResp = new(api.BaseResp)
	resp.BaseResp.Code = 0
	resp.BaseResp.Message = "success"

	c.JSON(pconsts.StatusOK, resp)
}
