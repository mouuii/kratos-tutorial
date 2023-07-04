package server

import (
	v1 "devops-agent/api/agent/v1"
	"devops-agent/internal/conf"
	"devops-agent/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// SelfErrorEncoder encodes the error to the HTTP response.
//func SelfErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
//	se := errors.FromError(err)
//	codec, _ := http.CodecForRequest(r, "Accept")
//
//	type SelfErrResp struct {
//		Code int `json:"errCode"`
//	}
//	e := SelfErrResp{
//		int(se.Code),
//	}
//
//	body, _ := codec.Marshal(e)
//
//	w.Header().Set("Content-Type", ContentType(codec.Name()))
//	w.WriteHeader(int(se.Code))
//	_, _ = w.Write(body)
//}
//
//// ContentType returns the content-type with base prefix.
//func ContentType(subtype string) string {
//	return strings.Join([]string{"application", subtype}, "/")
//}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, agent *service.AgentService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
		//http.ErrorEncoder(SelfErrorEncoder),
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
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterAgentHTTPServer(srv, agent)
	return srv
}
