package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Authorizer.
// Experimental.
type IAuthorizerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAuthorizerRef
type jsiiProxy_IAuthorizerRef struct {
	internal.Type__constructsIConstruct
}

