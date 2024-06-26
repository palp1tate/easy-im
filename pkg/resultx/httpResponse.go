package resultx

import (
	"context"
	"net/http"
	"reflect"

	"github.com/palp1tate/easy-im/pkg/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *Response {
	res := &Response{
		Code: 200,
		Msg:  "Success!",
	}
	if !reflect.ValueOf(data).IsNil() {
		res.Data = data
	}
	return res
}

func Fail(code int, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
	}
}

func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		errcode := errorx.ServerCommonError
		errmsg := errorx.ErrMsg(errcode)

		causeErr := errors.Cause(err)
		var e *zerr.CodeMsg
		if errors.As(causeErr, &e) {
			errcode = e.Code
			errmsg = e.Msg
		} else {
			if gStatus, ok := status.FromError(causeErr); ok {
				errcode = int(gStatus.Code())
				errmsg = gStatus.Message()
			}
		}

		// 日志记录
		logx.WithContext(ctx).Errorf("【%s】 err %v", name, err)

		return http.StatusBadRequest, Fail(errcode, errmsg)
	}
}
