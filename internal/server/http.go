package server

import (
	v1 "ai-meeting-server/api/template/v1"
	"ai-meeting-server/internal/conf"
	"ai-meeting-server/internal/service"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, dataConf *conf.Data, roleTemplate *service.RoleTemplateService, meetingTemplate *service.MeetingTemplateService, image *service.ImageService, meeting *service.MeetingService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRoleHTTPServer(srv, roleTemplate)
	v1.RegisterMeetingHTTPServer(srv, meetingTemplate)
	route := srv.Route("/")
	// resource api
	resource := route.Group("/api/resource")
	resource.POST("/uploadimage", image.UploadFile)
	// assets
	srv.HandlePrefix("/image", nethttp.FileServer(nethttp.Dir(dataConf.AssetsPath)))

	// meeting api
	meetingGroup := route.Group("/api/meeting")
	meetingGroup.POST("/newmeeting", meeting.NewMeeting)
	meetingGroup.GET("/history", meeting.History)
	meetingGroup.GET("/progress", meeting.Progress)
	meetingGroup.PUT("/start", meeting.Start)
	meetingGroup.PUT("/end", meeting.End)
	meetingGroup.GET("/getInputData", meeting.GetInputData)
	meetingGroup.POST("/input", meeting.Input)
	meetingGroup.PUT("/submitInput", meeting.SubmitInput)
	meetingGroup.GET("/getThinkingData", meeting.GetThinkingData)
	meetingGroup.POST("/saveThinkingPresets", meeting.SaveThinkingPresets)
	meetingGroup.POST("/thinkAndQuiz", meeting.ThinkAndQuiz)
	meetingGroup.POST("/thinkAndAnswer", meeting.ThinkAndAnswer)
	meetingGroup.POST("/thinkAndSummary", meeting.ThinkAndSummary)
	meetingGroup.GET("/getDiscussionData", meeting.GetDiscussionData)
	meetingGroup.POST("/saveDiscussionPresets", meeting.SaveDiscussionPresets)
	meetingGroup.POST("/discuss", meeting.Discuss)
	meetingGroup.POST("/discussAndQuiz", meeting.DiscussAndQuiz)
	meetingGroup.POST("/discussAndSummary", meeting.DiscussAndSummary)
	meetingGroup.GET("/getProcessingData", meeting.GetProcessingData)
	meetingGroup.POST("/process", meeting.Process)
	meetingGroup.GET("/getOutputData", meeting.GetOutputData)
	meetingGroup.POST("/output", meeting.Output)
	meetingGroup.POST("/chat", meeting.Chat)
	meetingGroup.OPTIONS("/{meeting}", meeting.MeetingOptions)

	return srv
}
