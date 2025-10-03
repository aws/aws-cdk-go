package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceEnvironment.
// Experimental.
type IServiceEnvironmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceEnvironmentRef
type jsiiProxy_IServiceEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

