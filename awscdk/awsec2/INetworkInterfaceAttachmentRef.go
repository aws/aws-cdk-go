package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterfaceAttachment.
// Experimental.
type INetworkInterfaceAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NetworkInterfaceAttachment resource.
	// Experimental.
	NetworkInterfaceAttachmentRef() *NetworkInterfaceAttachmentReference
}

// The jsii proxy for INetworkInterfaceAttachmentRef
type jsiiProxy_INetworkInterfaceAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INetworkInterfaceAttachmentRef) NetworkInterfaceAttachmentRef() *NetworkInterfaceAttachmentReference {
	var returns *NetworkInterfaceAttachmentReference
	_jsii_.Get(
		j,
		"networkInterfaceAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInterfaceAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInterfaceAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

