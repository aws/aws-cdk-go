package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ADMChannel.
// Experimental.
type IADMChannelRef interface {
	constructs.IConstruct
	// A reference to a ADMChannel resource.
	// Experimental.
	AdmChannelRef() *ADMChannelReference
}

// The jsii proxy for IADMChannelRef
type jsiiProxy_IADMChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IADMChannelRef) AdmChannelRef() *ADMChannelReference {
	var returns *ADMChannelReference
	_jsii_.Get(
		j,
		"admChannelRef",
		&returns,
	)
	return returns
}

