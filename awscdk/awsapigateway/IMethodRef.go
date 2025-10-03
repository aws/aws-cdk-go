package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Method.
// Experimental.
type IMethodRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMethodRef
type jsiiProxy_IMethodRef struct {
	internal.Type__constructsIConstruct
}

