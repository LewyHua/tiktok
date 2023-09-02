// Code generated by hertz generator.

package main

import (
	"crypto/tls"
	"crypto/x509"
	"douyin/config"
	"douyin/constant"
	"douyin/logger"
	"douyin/mw/redis"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/http2"
	"github.com/hertz-contrib/http2/factory"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	//_ "net/http/pprof"
)

func main() {
	// 加载配置
	if err := config.Init(); err != nil {
		zap.L().Error("Load config failed, err:%v\n", zap.Error(err))
		return
	}
	// 加载日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("Init logger failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化中间件: redis
	if err := redis.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}

	cert, err := tls.LoadX509KeyPair("./tls/www.godreamcode.top.pem", "./tls/www.godreamcode.top.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	certBytes, err := ioutil.ReadFile("./tls/Digicert G2 ROOT.cer")
	if err != nil {
		log.Fatal(err.Error())
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("Failed to parse root certificate.")
	}
	cfg := &tls.Config{
		Certificates:     []tls.Certificate{cert},
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		ClientCAs:  caCertPool,
		NextProtos: []string{http2.NextProtoTLS},
	}

	h := server.Default(
		server.WithHostPorts(constant.ApiServicePort),
		server.WithMaxRequestBodySize(50*1024*1024),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
		server.WithTLS(cfg),
		server.WithALPN(true),
	)
	h.AddProtocol("h2", factory.NewServerFactory())
	//pprof.Register(h)
	//go func() {
	//	ip := "0.0.0.0:8888"
	//	if err := http.ListenAndServe(ip, nil); err != nil {
	//		log.Printf("start pprof failed on %s\n", ip)
	//	}
	//}()

	customizedRegister(h)
	h.Spin()
}

func mainWithOutTLS() {
	// 加载配置
	if err := config.Init(); err != nil {
		zap.L().Error("Load config failed, err:%v\n", zap.Error(err))
		return
	}
	// 加载日志
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("Init logger failed, err:%v\n", zap.Error(err))
		return
	}

	// 初始化中间件: redis
	if err := redis.Init(config.Conf); err != nil {
		zap.L().Error("Init redis failed, err:%v\n", zap.Error(err))
		return
	}
	h := server.Default(
		server.WithHostPorts(constant.ApiServicePort),
		server.WithMaxRequestBodySize(50*1024*1024),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
	)
	//pprof.Register(h)
	//go func() {
	//	ip := "0.0.0.0:8888"
	//	if err := http.ListenAndServe(ip, nil); err != nil {
	//		log.Printf("start pprof failed on %s\n", ip)
	//	}
	//}()

	customizedRegister(h)
	h.Spin()
}
