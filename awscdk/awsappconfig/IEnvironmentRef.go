package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Environment.
// Experimental.
type IEnvironmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEnvironmentRef
type jsiiProxy_IEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

