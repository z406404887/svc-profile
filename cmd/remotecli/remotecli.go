// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/profile.proto
package remotecli

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	gomeetContext "github.com/gomeet/gomeet/utils/context"

	"github.com/gomeet-examples/svc-profile/client"
	pb "github.com/gomeet-examples/svc-profile/pb"
)

type ContextCall int

const (
	ConsoleCall ContextCall = iota // 0
	CliCall                        // 1
)

type RemoteCli interface {
	GetActionsMap() map[string]func([]string) (string, error)
	Eval(line string) (string, error)
	Close()
	RemoteVersion() (*pb.VersionResponse, error)
}

type remoteCli struct {
	name    string
	version string
	ctxCall ContextCall
	*client.GomeetClient
	c   pb.ProfileClient
	ctx context.Context
}

func NewRemoteCli(
	name string,
	version string,
	ctxCall ContextCall,
	addr string,
	timeout int,
	caCertificate string,
	clientCertificate string,
	clientPrivateKey string,
	jwt string,
) (RemoteCli, error) {
	cli, err := client.NewGomeetClient(
		addr,
		timeout,
		caCertificate,
		clientCertificate,
		clientPrivateKey,
	)

	if err != nil {
		return nil, err
	}

	// prepare the context
	ctx := gomeetContext.AuthContextFromJWT(context.Background(), jwt)

	rc := &remoteCli{
		name,
		version,
		ctxCall,
		cli,
		cli.GetGRPCClient(),
		ctx,
	}

	return rc, nil
}

// Close : close grpc connection `cc`
func (c *remoteCli) Close() {
	c.GomeetClient.Close()
}

// GetGRPCClient : getter on internal gRPC client
func (c *remoteCli) GetGRPCClient() pb.ProfileClient {
	return c.c
}

// SetNewJwt : Add update context with authorization header from jwt
func (c *remoteCli) SetJWT(jwt string) {
	c.ctx = gomeetContext.AuthContextFromJWT(context.Background(), jwt)
}

// UnsetNewJwt : unset authorization header from context
func (c *remoteCli) UnsetJWT() {
	c.ctx = context.Background()
}

func (c *remoteCli) RemoteVersion() (*pb.VersionResponse, error) {
	m := &pb.EmptyMessage{}
	r, err := c.c.Version(c.ctx, m)
	if err != nil {
		return nil, fmt.Errorf("Version service call fail - %v", err)
	}

	return r, nil
}

func (c *remoteCli) GetActionsMap() map[string]func([]string) (string, error) {
	return map[string]func([]string) (string, error){
		//grpc method
		"version":         c.cmdVersion,
		"services_status": c.cmdServicesStatus,
		"create":          c.cmdCreate,
		"read":            c.cmdRead,
		"list":            c.cmdList,
		"update":          c.cmdUpdate,
		"soft_delete":     c.cmdSoftDelete,
		"hard_delete":     c.cmdHardDelete,
		// global method
		"console_version": c.cmdConsoleVersion,
		"jwt":             c.cmdJWT,
		"service_address": c.cmdServiceAddress,
		"tls_config":      c.cmdTLSConfig,
		"help":            c.cmdHelp,
	}
}

// Eval string to gRPC call
func (c *remoteCli) Eval(line string) (string, error) {
	if len(line) == 0 {
		return "", errors.New("Bad arguments : Empty")
	}

	for k, fn := range c.GetActionsMap() {
		if strings.HasPrefix(line, k) {
			l := len(k)
			params := strings.TrimSpace(line[l:])
			args := parseParams(params)
			return fn(args)
		}
	}

	return "", fmt.Errorf("Bad arguments : %v unknow", strconv.Quote(line))
}
