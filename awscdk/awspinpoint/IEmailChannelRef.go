package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailChannel.
// Experimental.
type IEmailChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EmailChannel resource.
	// Experimental.
	EmailChannelRef() *EmailChannelReference
}

// The jsii proxy for IEmailChannelRef
type jsiiProxy_IEmailChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEmailChannelRef) EmailChannelRef() *EmailChannelReference {
	var returns *EmailChannelReference
	_jsii_.Get(
		j,
		"emailChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

