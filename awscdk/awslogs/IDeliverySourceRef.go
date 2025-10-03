package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliverySource.
// Experimental.
type IDeliverySourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeliverySourceRef
type jsiiProxy_IDeliverySourceRef struct {
	internal.Type__constructsIConstruct
}

