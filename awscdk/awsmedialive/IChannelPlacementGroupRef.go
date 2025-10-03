package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelPlacementGroup.
// Experimental.
type IChannelPlacementGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChannelPlacementGroupRef
type jsiiProxy_IChannelPlacementGroupRef struct {
	internal.Type__constructsIConstruct
}

