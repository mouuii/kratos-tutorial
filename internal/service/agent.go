package service

import (
	"context"
	"devops-agent/internal/biz"

	pb "devops-agent/api/agent/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer
}

func NewAgentService(uc *biz.GreeterUsecase) *AgentService {
	return &AgentService{}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentReply, error) {
	return &pb.CreateAgentReply{
		Result: "hello" + req.Name,
	}, nil
}
func (s *AgentService) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentReply, error) {
	return &pb.UpdateAgentReply{}, nil
}
func (s *AgentService) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentReply, error) {
	return &pb.DeleteAgentReply{}, nil
}
func (s *AgentService) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentReply, error) {
	return &pb.GetAgentReply{}, nil
}
func (s *AgentService) ListAgent(ctx context.Context, req *pb.ListAgentRequest) (*pb.ListAgentReply, error) {
	return &pb.ListAgentReply{}, nil
}
