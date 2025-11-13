package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCBlockPublicAccessOptions.
// Experimental.
type IVPCBlockPublicAccessOptionsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VPCBlockPublicAccessOptions resource.
	// Experimental.
	VpcBlockPublicAccessOptionsRef() *VPCBlockPublicAccessOptionsReference
}

// The jsii proxy for IVPCBlockPublicAccessOptionsRef
type jsiiProxy_IVPCBlockPublicAccessOptionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVPCBlockPublicAccessOptionsRef) VpcBlockPublicAccessOptionsRef() *VPCBlockPublicAccessOptionsReference {
	var returns *VPCBlockPublicAccessOptionsReference
	_jsii_.Get(
		j,
		"vpcBlockPublicAccessOptionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCBlockPublicAccessOptionsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCBlockPublicAccessOptionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

