package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSVoipChannel.
// Experimental.
type IAPNSVoipChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAPNSVoipChannelRef
type jsiiProxy_IAPNSVoipChannelRef struct {
	internal.Type__constructsIConstruct
}

