package service

import (
	pb "ai-meeting-server/api/template/v1"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRoleTemplateService, NewMeetingTemplateService, NewImageService, NewMeetingService)

// todo 暂时简单处理，后续在完善code状态
func RespBool(ok bool, err error) (*pb.BoolReply, error) {
	if err != nil {
		return &pb.BoolReply{
			Code: 400,
			Msg:  err.Error(),
			Data: &pb.BoolReply_Data{Ok: ok},
		}, err
	} else {
		return &pb.BoolReply{
			Code: 200,
			Msg:  "success",
			Data: &pb.BoolReply_Data{Ok: ok},
		}, nil
	}
}
