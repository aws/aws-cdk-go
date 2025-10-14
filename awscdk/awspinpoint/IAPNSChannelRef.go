package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSChannel.
// Experimental.
type IAPNSChannelRef interface {
	constructs.IConstruct
	// A reference to a APNSChannel resource.
	// Experimental.
	ApnsChannelRef() *APNSChannelReference
}

// The jsii proxy for IAPNSChannelRef
type jsiiProxy_IAPNSChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAPNSChannelRef) ApnsChannelRef() *APNSChannelReference {
	var returns *APNSChannelReference
	_jsii_.Get(
		j,
		"apnsChannelRef",
		&returns,
	)
	return returns
}

