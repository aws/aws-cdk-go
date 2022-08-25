// The CDK Construct Library for AWS::IVS
package awscdkivsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkivsalpha/v2/internal"
)

// Represents an IVS Channel.
// Experimental.
type IChannel interface {
	awscdk.IResource
	// Adds a stream key for this IVS Channel.
	// Experimental.
	AddStreamKey(id *string) StreamKey
	// The channel ARN.
	//
	// For example: arn:aws:ivs:us-west-2:123456789012:channel/abcdABCDefgh.
	// Experimental.
	ChannelArn() *string
}

// The jsii proxy for IChannel
type jsiiProxy_IChannel struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannel) AddStreamKey(id *string) StreamKey {
	var returns StreamKey

	_jsii_.Invoke(
		i,
		"addStreamKey",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IChannel) ChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelArn",
		&returns,
	)
	return returns
}

