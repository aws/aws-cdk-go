package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailChannel.
// Experimental.
type IEmailChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEmailChannelRef
type jsiiProxy_IEmailChannelRef struct {
	internal.Type__constructsIConstruct
}

