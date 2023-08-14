// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteservice

import (
	"context"
	favorite "douyin/kitex_gen/favorite"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":        kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"GetFavoriteList":       kitex.NewMethodInfo(getFavoriteListHandler, newGetFavoriteListArgs, newGetFavoriteListResult, false),
		"GetVideoFavoriteCount": kitex.NewMethodInfo(getVideoFavoriteCountHandler, newGetVideoFavoriteCountArgs, newGetVideoFavoriteCountResult, false),
		"GetUserFavoriteCount":  kitex.NewMethodInfo(getUserFavoriteCountHandler, newGetUserFavoriteCountArgs, newGetUserFavoriteCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(favorite.FavoriteService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *favorite.FavoriteActionRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *favorite.FavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *favorite.FavoriteActionRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *favorite.FavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *favorite.FavoriteActionResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *favorite.FavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).GetFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteListArgs:
		success, err := handler.(favorite.FavoriteService).GetFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteListArgs() interface{} {
	return &GetFavoriteListArgs{}
}

func newGetFavoriteListResult() interface{} {
	return &GetFavoriteListResult{}
}

type GetFavoriteListArgs struct {
	Req *favorite.FavoriteListRequest
}

func (p *GetFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteListArgs_Req_DEFAULT *favorite.FavoriteListRequest

func (p *GetFavoriteListArgs) GetReq() *favorite.FavoriteListRequest {
	if !p.IsSetReq() {
		return GetFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFavoriteListResult struct {
	Success *favorite.FavoriteListResponse
}

var GetFavoriteListResult_Success_DEFAULT *favorite.FavoriteListResponse

func (p *GetFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteListResult) GetSuccess() *favorite.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteListResponse)
}

func (p *GetFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFavoriteListResult) GetResult() interface{} {
	return p.Success
}

func getVideoFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.VideoFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).GetVideoFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideoFavoriteCountArgs:
		success, err := handler.(favorite.FavoriteService).GetVideoFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideoFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideoFavoriteCountArgs() interface{} {
	return &GetVideoFavoriteCountArgs{}
}

func newGetVideoFavoriteCountResult() interface{} {
	return &GetVideoFavoriteCountResult{}
}

type GetVideoFavoriteCountArgs struct {
	Req *favorite.VideoFavoriteCountRequest
}

func (p *GetVideoFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.VideoFavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetVideoFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetVideoFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetVideoFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideoFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(favorite.VideoFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideoFavoriteCountArgs_Req_DEFAULT *favorite.VideoFavoriteCountRequest

func (p *GetVideoFavoriteCountArgs) GetReq() *favorite.VideoFavoriteCountRequest {
	if !p.IsSetReq() {
		return GetVideoFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetVideoFavoriteCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetVideoFavoriteCountResult struct {
	Success *favorite.VideoFavoriteCountResponse
}

var GetVideoFavoriteCountResult_Success_DEFAULT *favorite.VideoFavoriteCountResponse

func (p *GetVideoFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.VideoFavoriteCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetVideoFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetVideoFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetVideoFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideoFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(favorite.VideoFavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoFavoriteCountResult) GetSuccess() *favorite.VideoFavoriteCountResponse {
	if !p.IsSetSuccess() {
		return GetVideoFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.VideoFavoriteCountResponse)
}

func (p *GetVideoFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetVideoFavoriteCountResult) GetResult() interface{} {
	return p.Success
}

func getUserFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.UserFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).GetUserFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserFavoriteCountArgs:
		success, err := handler.(favorite.FavoriteService).GetUserFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserFavoriteCountArgs() interface{} {
	return &GetUserFavoriteCountArgs{}
}

func newGetUserFavoriteCountResult() interface{} {
	return &GetUserFavoriteCountResult{}
}

type GetUserFavoriteCountArgs struct {
	Req *favorite.UserFavoriteCountRequest
}

func (p *GetUserFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.UserFavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(favorite.UserFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserFavoriteCountArgs_Req_DEFAULT *favorite.UserFavoriteCountRequest

func (p *GetUserFavoriteCountArgs) GetReq() *favorite.UserFavoriteCountRequest {
	if !p.IsSetReq() {
		return GetUserFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserFavoriteCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserFavoriteCountResult struct {
	Success *favorite.UserFavoriteCountResponse
}

var GetUserFavoriteCountResult_Success_DEFAULT *favorite.UserFavoriteCountResponse

func (p *GetUserFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.UserFavoriteCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(favorite.UserFavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserFavoriteCountResult) GetSuccess() *favorite.UserFavoriteCountResponse {
	if !p.IsSetSuccess() {
		return GetUserFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.UserFavoriteCountResponse)
}

func (p *GetUserFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserFavoriteCountResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *favorite.FavoriteActionRequest) (r *favorite.FavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, Req *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	var _args GetFavoriteListArgs
	_args.Req = Req
	var _result GetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoFavoriteCount(ctx context.Context, Req *favorite.VideoFavoriteCountRequest) (r *favorite.VideoFavoriteCountResponse, err error) {
	var _args GetVideoFavoriteCountArgs
	_args.Req = Req
	var _result GetVideoFavoriteCountResult
	if err = p.c.Call(ctx, "GetVideoFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFavoriteCount(ctx context.Context, Req *favorite.UserFavoriteCountRequest) (r *favorite.UserFavoriteCountResponse, err error) {
	var _args GetUserFavoriteCountArgs
	_args.Req = Req
	var _result GetUserFavoriteCountResult
	if err = p.c.Call(ctx, "GetUserFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
