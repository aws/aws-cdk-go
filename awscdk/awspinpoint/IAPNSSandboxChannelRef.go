package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSSandboxChannel.
// Experimental.
type IAPNSSandboxChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAPNSSandboxChannelRef
type jsiiProxy_IAPNSSandboxChannelRef struct {
	internal.Type__constructsIConstruct
}

