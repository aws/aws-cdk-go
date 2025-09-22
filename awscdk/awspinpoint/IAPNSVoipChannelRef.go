package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSVoipChannel.
// Experimental.
type IAPNSVoipChannelRef interface {
	constructs.IConstruct
	// A reference to a APNSVoipChannel resource.
	// Experimental.
	ApnsVoipChannelRef() *APNSVoipChannelReference
}

// The jsii proxy for IAPNSVoipChannelRef
type jsiiProxy_IAPNSVoipChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAPNSVoipChannelRef) ApnsVoipChannelRef() *APNSVoipChannelReference {
	var returns *APNSVoipChannelReference
	_jsii_.Get(
		j,
		"apnsVoipChannelRef",
		&returns,
	)
	return returns
}

