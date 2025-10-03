package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ChannelNamespace.
// Experimental.
type IChannelNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IChannelNamespaceRef
type jsiiProxy_IChannelNamespaceRef struct {
	internal.Type__constructsIConstruct
}

