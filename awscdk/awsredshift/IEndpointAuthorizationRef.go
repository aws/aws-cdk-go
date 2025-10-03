package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EndpointAuthorization.
// Experimental.
type IEndpointAuthorizationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEndpointAuthorizationRef
type jsiiProxy_IEndpointAuthorizationRef struct {
	internal.Type__constructsIConstruct
}

