package awsapplicationsignals

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLevelObjective.
// Experimental.
type IServiceLevelObjectiveRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceLevelObjectiveRef
type jsiiProxy_IServiceLevelObjectiveRef struct {
	internal.Type__constructsIConstruct
}

