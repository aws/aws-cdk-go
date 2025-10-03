package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Api.
// Experimental.
type IApiRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApiRef
type jsiiProxy_IApiRef struct {
	internal.Type__constructsIConstruct
}

