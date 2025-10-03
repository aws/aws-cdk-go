package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ADMChannel.
// Experimental.
type IADMChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IADMChannelRef
type jsiiProxy_IADMChannelRef struct {
	internal.Type__constructsIConstruct
}

