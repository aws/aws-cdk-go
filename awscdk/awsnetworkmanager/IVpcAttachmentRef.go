package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcAttachment.
// Experimental.
type IVpcAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VpcAttachment resource.
	// Experimental.
	VpcAttachmentRef() *VpcAttachmentReference
}

// The jsii proxy for IVpcAttachmentRef
type jsiiProxy_IVpcAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVpcAttachmentRef) VpcAttachmentRef() *VpcAttachmentReference {
	var returns *VpcAttachmentReference
	_jsii_.Get(
		j,
		"vpcAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

