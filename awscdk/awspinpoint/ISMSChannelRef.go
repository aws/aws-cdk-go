package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SMSChannel.
// Experimental.
type ISMSChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SMSChannel resource.
	// Experimental.
	SmsChannelRef() *SMSChannelReference
}

// The jsii proxy for ISMSChannelRef
type jsiiProxy_ISMSChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISMSChannelRef) SmsChannelRef() *SMSChannelReference {
	var returns *SMSChannelReference
	_jsii_.Get(
		j,
		"smsChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISMSChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISMSChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

