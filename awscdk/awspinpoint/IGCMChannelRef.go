package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GCMChannel.
// Experimental.
type IGCMChannelRef interface {
	constructs.IConstruct
	// A reference to a GCMChannel resource.
	// Experimental.
	GcmChannelRef() *GCMChannelReference
}

// The jsii proxy for IGCMChannelRef
type jsiiProxy_IGCMChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGCMChannelRef) GcmChannelRef() *GCMChannelReference {
	var returns *GCMChannelReference
	_jsii_.Get(
		j,
		"gcmChannelRef",
		&returns,
	)
	return returns
}

