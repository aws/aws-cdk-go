package awsmediapackagev2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPolicy.
// Experimental.
type IChannelPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChannelPolicyRef
type jsiiProxy_IChannelPolicyRef struct {
	internal.Type__constructsIConstruct
}

