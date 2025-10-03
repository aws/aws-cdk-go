package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiMapping.
// Experimental.
type IApiMappingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApiMappingRef
type jsiiProxy_IApiMappingRef struct {
	internal.Type__constructsIConstruct
}

