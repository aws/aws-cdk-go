package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConsumableResource.
// Experimental.
type IConsumableResourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConsumableResourceRef
type jsiiProxy_IConsumableResourceRef struct {
	internal.Type__constructsIConstruct
}

