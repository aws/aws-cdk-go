package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionConfiguration.
// Experimental.
type IFunctionConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFunctionConfigurationRef
type jsiiProxy_IFunctionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

