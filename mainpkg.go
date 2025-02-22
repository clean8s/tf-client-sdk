package tfclient

import (
	goplug "github.com/hashicorp/go-plugin"
	"os/exec"
	"github.com/clean8s/tf-client-sdk/configschema"
	"github.com/clean8s/tf-client-sdk/plugin"
	"github.com/clean8s/tf-client-sdk/providers"
)

type Block = configschema.Block
type Attribute = configschema.Attribute
type NestedBlock = configschema.NestedBlock
type Object = configschema.Object
type ValidateProviderConfigRequest = providers.ValidateProviderConfigRequest
type ValidateResourceConfigRequest = providers.ValidateResourceConfigRequest
type ConfigureProviderRequest = providers.ConfigureProviderRequest
type PlanResourceChangeRequest = providers.PlanResourceChangeRequest
type ApplyResourceChangeRequest = providers.ApplyResourceChangeRequest
type GRPCProvider = plugin.GRPCProvider

var TF_VERSION = Version

// MakeClient creates a Terraform client from a command and/or a RPC port. It's used to access
// a provider in a clean way.
//   c := MakeClient(exec.Command("tf-provider-example.exe"), nil)
//   c.ReadResource(...)
func MakeClient(cmd *exec.Cmd, rc *goplug.ReattachConfig) *plugin.GRPCProvider {
	config := &goplug.ClientConfig{
		Reattach:         rc,
		HandshakeConfig:  plugin.Handshake,
		AllowedProtocols: []goplug.Protocol{goplug.ProtocolGRPC},
		Managed:          true,
		Cmd:              cmd,
		AutoMTLS:         false,
		VersionedPlugins: plugin.VersionedPlugins,
		Plugins:          plugin.VersionedPlugins[5],
	}

	client := goplug.NewClient(config)
	rpcClient, err := client.Client()
	_ = err
	raw, err := rpcClient.Dispense("provider")
	p := raw.(*plugin.GRPCProvider)
	p.PluginClient = client
	return p
}

// MakeClientCustom acts like MakeClient except you can pass
// your custom ClientConfig. Use MakeClientConfig() for some
// defaults.
func MakeClientCustom(config *goplug.ClientConfig) *plugin.GRPCProvider {
	client := goplug.NewClient(config)
	rpcClient, err := client.Client()
	_ = err
	raw, err := rpcClient.Dispense("provider")
	p := raw.(*plugin.GRPCProvider)
	p.PluginClient = client
	return p
}

// MakeClientConfig creates a go-plugin ClientConfig object optimal
// for connecting to a provider.
func MakeClientConfig() *goplug.ClientConfig {
	config := &goplug.ClientConfig{
		HandshakeConfig:  plugin.Handshake,
		AllowedProtocols: []goplug.Protocol{goplug.ProtocolGRPC},
		Managed:          true,
		AutoMTLS:         false,
		VersionedPlugins: plugin.VersionedPlugins,
		Plugins:          plugin.VersionedPlugins[5],
	}
	return config
}
