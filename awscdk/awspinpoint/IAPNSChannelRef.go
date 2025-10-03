package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSChannel.
// Experimental.
type IAPNSChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAPNSChannelRef
type jsiiProxy_IAPNSChannelRef struct {
	internal.Type__constructsIConstruct
}

