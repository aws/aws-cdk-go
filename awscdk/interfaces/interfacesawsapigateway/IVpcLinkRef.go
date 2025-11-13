package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcLink.
// Experimental.
type IVpcLinkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VpcLink resource.
	// Experimental.
	VpcLinkRef() *VpcLinkReference
}

// The jsii proxy for IVpcLinkRef
type jsiiProxy_IVpcLinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVpcLinkRef) VpcLinkRef() *VpcLinkReference {
	var returns *VpcLinkReference
	_jsii_.Get(
		j,
		"vpcLinkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcLinkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcLinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

