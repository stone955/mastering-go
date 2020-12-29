package options

import (
	"crypto/tls"
	"time"
)

// Server 服务端结构体
type Server struct {
	Addr        string        // 必填项
	Port        int           // 必填项
	Protocol    string        // 选填 默认 tcp
	DialTimeout time.Duration // 选填 默认 30s
	MaxConn     int           // 选填 默认 100
	TLS         *tls.Config   // 选填 默认 nil
}

// NewServer 创建 Server
func NewServer(addr string, port int, options ...Option) (*Server, error) {
	s := &Server{
		Addr:        addr,
		Port:        port,
		Protocol:    "tcp",
		DialTimeout: 30 * time.Second,
		MaxConn:     100,
		TLS:         nil,
	}

	for _, option := range options {
		option(s)
	}

	// ...
	return s, nil
}

// Option 可选参数配置函数
type Option func(s *Server)

// Protocol 配置协议类型
func Protocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

// DialTimeout 配置连接超时
func DialTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.DialTimeout = timeout
	}
}

// MaxConn 配置最大连接数
func MaxConn(maxConn int) Option {
	return func(s *Server) {
		s.MaxConn = maxConn
	}
}

// TLS 配置TLS
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}
