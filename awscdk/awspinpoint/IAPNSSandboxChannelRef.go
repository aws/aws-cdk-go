package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSSandboxChannel.
// Experimental.
type IAPNSSandboxChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a APNSSandboxChannel resource.
	// Experimental.
	ApnsSandboxChannelRef() *APNSSandboxChannelReference
}

// The jsii proxy for IAPNSSandboxChannelRef
type jsiiProxy_IAPNSSandboxChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAPNSSandboxChannelRef) ApnsSandboxChannelRef() *APNSSandboxChannelReference {
	var returns *APNSSandboxChannelReference
	_jsii_.Get(
		j,
		"apnsSandboxChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSSandboxChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSSandboxChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

