package interfacesawsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEConfiguration.
// Experimental.
type IVPCEConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VPCEConfiguration resource.
	// Experimental.
	VpceConfigurationRef() *VPCEConfigurationReference
}

// The jsii proxy for IVPCEConfigurationRef
type jsiiProxy_IVPCEConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVPCEConfigurationRef) VpceConfigurationRef() *VPCEConfigurationReference {
	var returns *VPCEConfigurationReference
	_jsii_.Get(
		j,
		"vpceConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCEConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCEConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

