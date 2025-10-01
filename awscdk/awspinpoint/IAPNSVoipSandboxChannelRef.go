package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSVoipSandboxChannel.
// Experimental.
type IAPNSVoipSandboxChannelRef interface {
	constructs.IConstruct
	// A reference to a APNSVoipSandboxChannel resource.
	// Experimental.
	ApnsVoipSandboxChannelRef() *APNSVoipSandboxChannelReference
}

// The jsii proxy for IAPNSVoipSandboxChannelRef
type jsiiProxy_IAPNSVoipSandboxChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAPNSVoipSandboxChannelRef) ApnsVoipSandboxChannelRef() *APNSVoipSandboxChannelReference {
	var returns *APNSVoipSandboxChannelReference
	_jsii_.Get(
		j,
		"apnsVoipSandboxChannelRef",
		&returns,
	)
	return returns
}

