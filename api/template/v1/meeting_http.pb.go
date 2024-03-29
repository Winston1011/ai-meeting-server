// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.19.4
// source: template/v1/meeting.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMeetingCreateMeeting = "/api.template.v1.Meeting/CreateMeeting"
const OperationMeetingDeleteMeeting = "/api.template.v1.Meeting/DeleteMeeting"
const OperationMeetingGetMeeting = "/api.template.v1.Meeting/GetMeeting"
const OperationMeetingListMeeting = "/api.template.v1.Meeting/ListMeeting"
const OperationMeetingUpdateMeeting = "/api.template.v1.Meeting/UpdateMeeting"

type MeetingHTTPServer interface {
	CreateMeeting(context.Context, *CreateMeetingRequest) (*BoolReply, error)
	DeleteMeeting(context.Context, *DeleteMeetingRequest) (*BoolReply, error)
	GetMeeting(context.Context, *GetMeetingRequest) (*GetMeetingReply, error)
	ListMeeting(context.Context, *ListMeetingRequest) (*ListMeetingReply, error)
	UpdateMeeting(context.Context, *UpdateMeetingRequest) (*BoolReply, error)
}

func RegisterMeetingHTTPServer(s *http.Server, srv MeetingHTTPServer) {
	r := s.Route("/")
	r.POST("/api/template/meeting", _Meeting_CreateMeeting0_HTTP_Handler(srv))
	r.PUT("/api/template/meeting", _Meeting_UpdateMeeting0_HTTP_Handler(srv))
	r.DELETE("/api/template/meeting", _Meeting_DeleteMeeting0_HTTP_Handler(srv))
	r.GET("/api/template/meeting/{id}", _Meeting_GetMeeting0_HTTP_Handler(srv))
	r.GET("/api/template/meeting", _Meeting_ListMeeting0_HTTP_Handler(srv))
}

func _Meeting_CreateMeeting0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMeetingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMeetingCreateMeeting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMeeting(ctx, req.(*CreateMeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BoolReply)
		return ctx.Result(200, reply)
	}
}

func _Meeting_UpdateMeeting0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMeetingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMeetingUpdateMeeting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMeeting(ctx, req.(*UpdateMeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BoolReply)
		return ctx.Result(200, reply)
	}
}

func _Meeting_DeleteMeeting0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMeetingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMeetingDeleteMeeting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMeeting(ctx, req.(*DeleteMeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BoolReply)
		return ctx.Result(200, reply)
	}
}

func _Meeting_GetMeeting0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMeetingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMeetingGetMeeting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMeeting(ctx, req.(*GetMeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMeetingReply)
		return ctx.Result(200, reply)
	}
}

func _Meeting_ListMeeting0_HTTP_Handler(srv MeetingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMeetingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMeetingListMeeting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMeeting(ctx, req.(*ListMeetingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMeetingReply)
		return ctx.Result(200, reply)
	}
}

type MeetingHTTPClient interface {
	CreateMeeting(ctx context.Context, req *CreateMeetingRequest, opts ...http.CallOption) (rsp *BoolReply, err error)
	DeleteMeeting(ctx context.Context, req *DeleteMeetingRequest, opts ...http.CallOption) (rsp *BoolReply, err error)
	GetMeeting(ctx context.Context, req *GetMeetingRequest, opts ...http.CallOption) (rsp *GetMeetingReply, err error)
	ListMeeting(ctx context.Context, req *ListMeetingRequest, opts ...http.CallOption) (rsp *ListMeetingReply, err error)
	UpdateMeeting(ctx context.Context, req *UpdateMeetingRequest, opts ...http.CallOption) (rsp *BoolReply, err error)
}

type MeetingHTTPClientImpl struct {
	cc *http.Client
}

func NewMeetingHTTPClient(client *http.Client) MeetingHTTPClient {
	return &MeetingHTTPClientImpl{client}
}

func (c *MeetingHTTPClientImpl) CreateMeeting(ctx context.Context, in *CreateMeetingRequest, opts ...http.CallOption) (*BoolReply, error) {
	var out BoolReply
	pattern := "/api/template/meeting"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMeetingCreateMeeting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MeetingHTTPClientImpl) DeleteMeeting(ctx context.Context, in *DeleteMeetingRequest, opts ...http.CallOption) (*BoolReply, error) {
	var out BoolReply
	pattern := "/api/template/meeting"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMeetingDeleteMeeting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MeetingHTTPClientImpl) GetMeeting(ctx context.Context, in *GetMeetingRequest, opts ...http.CallOption) (*GetMeetingReply, error) {
	var out GetMeetingReply
	pattern := "/api/template/meeting/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMeetingGetMeeting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MeetingHTTPClientImpl) ListMeeting(ctx context.Context, in *ListMeetingRequest, opts ...http.CallOption) (*ListMeetingReply, error) {
	var out ListMeetingReply
	pattern := "/api/template/meeting"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMeetingListMeeting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MeetingHTTPClientImpl) UpdateMeeting(ctx context.Context, in *UpdateMeetingRequest, opts ...http.CallOption) (*BoolReply, error) {
	var out BoolReply
	pattern := "/api/template/meeting"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMeetingUpdateMeeting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
