package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectAttachment.
// Experimental.
type IConnectAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConnectAttachment resource.
	// Experimental.
	ConnectAttachmentRef() *ConnectAttachmentReference
}

// The jsii proxy for IConnectAttachmentRef
type jsiiProxy_IConnectAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConnectAttachmentRef) ConnectAttachmentRef() *ConnectAttachmentReference {
	var returns *ConnectAttachmentReference
	_jsii_.Get(
		j,
		"connectAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

