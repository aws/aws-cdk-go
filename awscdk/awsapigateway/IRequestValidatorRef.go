package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RequestValidator.
// Experimental.
type IRequestValidatorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRequestValidatorRef
type jsiiProxy_IRequestValidatorRef struct {
	internal.Type__constructsIConstruct
}

