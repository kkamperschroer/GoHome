package plugin_loader

import (
	"github.com/pivotal-golang/lager"
	"net"
	"net/http"
	"net/rpc"
)

type PluginServer interface {
	Start() error
	Test(string, *StringResponse) error
}

type PluginServerInstance struct {
	logger lager.Logger
}

type StringResponse struct {
	StringResponse string
}

func NewPluginServer(logger lager.Logger) PluginServer {
	return &PluginServerInstance{
		logger: logger,
	}
}

func (p *PluginServerInstance) Start() error {
	p.logger.Debug("Starting plugin server...")

	rpc.Register(p)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4249")
	if err != nil {
		p.logger.Error("Error starting RPC client", err)
		return err
	}

	go http.Serve(listener, nil)

	p.logger.Debug("Plugin server started")
	return nil
}

func (p *PluginServerInstance) Test(request string, response *StringResponse) error {
	p.logger.Debug("Plugin server received Test call: " + request)

	response.StringResponse = "Hello from GoHome!"
	return nil
}
