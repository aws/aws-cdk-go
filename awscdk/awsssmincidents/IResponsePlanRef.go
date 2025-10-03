package awsssmincidents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponsePlan.
// Experimental.
type IResponsePlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResponsePlanRef
type jsiiProxy_IResponsePlanRef struct {
	internal.Type__constructsIConstruct
}

