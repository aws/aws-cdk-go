package awsnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelAssociation.
// Experimental.
type IChannelAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChannelAssociationRef
type jsiiProxy_IChannelAssociationRef struct {
	internal.Type__constructsIConstruct
}

