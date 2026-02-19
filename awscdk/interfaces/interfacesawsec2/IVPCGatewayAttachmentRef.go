package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCGatewayAttachment.
// Experimental.
type IVPCGatewayAttachmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VPCGatewayAttachment resource.
	// Experimental.
	VpcGatewayAttachmentRef() *VPCGatewayAttachmentReference
}

// The jsii proxy for IVPCGatewayAttachmentRef
type jsiiProxy_IVPCGatewayAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVPCGatewayAttachmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVPCGatewayAttachmentRef) VpcGatewayAttachmentRef() *VPCGatewayAttachmentReference {
	var returns *VPCGatewayAttachmentReference
	_jsii_.Get(
		j,
		"vpcGatewayAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCGatewayAttachmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCGatewayAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

