package service

import (
	"context"
	"devops-agent/internal/biz"

	pb "devops-agent/api/agent/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer
	uc *biz.UserUsecase
}

func NewAgentService(uc *biz.UserUsecase) *AgentService {
	return &AgentService{
		uc: uc,
	}
}

func (s *AgentService) RegistryUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user, err := s.uc.Register(ctx, &biz.User{req.Name, req.Age})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Result: "hello" + user.Name,
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
