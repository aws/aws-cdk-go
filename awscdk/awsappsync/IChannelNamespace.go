package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
)

// An AppSync channel namespace.
type IChannelNamespace interface {
	awscdk.IResource
	// The ARN of the AppSync channel namespace.
	ChannelNamespaceArn() *string
}

// The jsii proxy for IChannelNamespace
type jsiiProxy_IChannelNamespace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IChannelNamespace) ChannelNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNamespaceArn",
		&returns,
	)
	return returns
}

