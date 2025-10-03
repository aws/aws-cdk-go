package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Delivery.
// Experimental.
type IDeliveryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeliveryRef
type jsiiProxy_IDeliveryRef struct {
	internal.Type__constructsIConstruct
}

