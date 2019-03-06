package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/dy-dayan/common-srv-atomicid/dal/db"
	"github.com/dy-dayan/common-srv-atomicid/idl"
	srv "github.com/dy-dayan/common-srv-atomicid/idl/dayan/common/srv-atomicid"
)

type Handler struct {
}

// 新建议案
func (h *Handler) GetID(ctx context.Context, req *srv.GetIDReq, rsp *srv.GetIDResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	if len(req.Label) == 0 {
		rsp.BaseResp.Code = int32(base.CODE_INVALID_PARAMETER)
		rsp.BaseResp.Msg = "invalid parameters"
		return nil
	}

	id, err := db.GetID(req.Label)
	if err != nil {
		logrus.Warnf("db.GetID error:%v", err)
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = err.Error()
		return nil
	}

	rsp.Id = id

	return nil
}
