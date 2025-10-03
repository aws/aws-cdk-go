package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GCMChannel.
// Experimental.
type IGCMChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGCMChannelRef
type jsiiProxy_IGCMChannelRef struct {
	internal.Type__constructsIConstruct
}

