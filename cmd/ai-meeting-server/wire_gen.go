// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ai-meeting-server/internal/biz"
	"ai-meeting-server/internal/conf"
	"ai-meeting-server/internal/data"
	"ai-meeting-server/internal/server"
	"ai-meeting-server/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(confData, logger, database)
	if err != nil {
		return nil, nil, err
	}
	roleTemplateRepo := data.NewRoleTemplateRepo(dataData, logger)
	roleTemplateUsecase := biz.NewRoleTemplateUsecase(roleTemplateRepo, logger)
	roleTemplateService := service.NewRoleTemplateService(roleTemplateUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, roleTemplateService, logger)
	meetingTemplateRepo := data.NewMeetingTemplateRepo(dataData, logger)
	meetingTemplateUsecase := biz.NewMeetingTemplateUsecase(meetingTemplateRepo, roleTemplateRepo, logger)
	meetingTemplateService := service.NewMeetingTemplateService(meetingTemplateUsecase, logger)
	imageUsecase := biz.NewImageUsecase(confData, logger)
	imageService := service.NewImageService(imageUsecase, confData, logger)
	meetingRepo := data.NewMeetingRepo(dataData, logger)
	gpt := biz.NewGpt(confData, logger)
	meetingUsecase := biz.NewMeetingUsecase(meetingRepo, meetingTemplateRepo, gpt, logger)
	meetingService := service.NewMeetingService(meetingUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, confData, roleTemplateService, meetingTemplateService, imageService, meetingService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
