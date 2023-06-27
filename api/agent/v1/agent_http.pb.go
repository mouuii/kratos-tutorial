// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.6
// source: agent/v1/agent.proto

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

const OperationAgentRegistryUser = "/api.agent.v1.Agent/RegistryUser"
const OperationAgentUpdateAgent = "/api.agent.v1.Agent/UpdateAgent"

type AgentHTTPServer interface {
	RegistryUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error)
}

func RegisterAgentHTTPServer(s *http.Server, srv AgentHTTPServer) {
	r := s.Route("/")
	r.POST("/user/register", _Agent_RegistryUser0_HTTP_Handler(srv))
	r.GET("/devops", _Agent_UpdateAgent0_HTTP_Handler(srv))
}

func _Agent_RegistryUser0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentRegistryUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegistryUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _Agent_UpdateAgent0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAgentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentUpdateAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAgent(ctx, req.(*UpdateAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAgentReply)
		return ctx.Result(200, reply)
	}
}

type AgentHTTPClient interface {
	RegistryUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	UpdateAgent(ctx context.Context, req *UpdateAgentRequest, opts ...http.CallOption) (rsp *UpdateAgentReply, err error)
}

type AgentHTTPClientImpl struct {
	cc *http.Client
}

func NewAgentHTTPClient(client *http.Client) AgentHTTPClient {
	return &AgentHTTPClientImpl{client}
}

func (c *AgentHTTPClientImpl) RegistryUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentRegistryUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...http.CallOption) (*UpdateAgentReply, error) {
	var out UpdateAgentReply
	pattern := "/devops"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentUpdateAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
