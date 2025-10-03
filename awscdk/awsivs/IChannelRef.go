package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Channel.
// Experimental.
type IChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChannelRef
type jsiiProxy_IChannelRef struct {
	internal.Type__constructsIConstruct
}

