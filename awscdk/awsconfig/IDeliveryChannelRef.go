package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryChannel.
// Experimental.
type IDeliveryChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeliveryChannelRef
type jsiiProxy_IDeliveryChannelRef struct {
	internal.Type__constructsIConstruct
}

