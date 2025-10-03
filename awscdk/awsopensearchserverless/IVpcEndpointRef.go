package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpoint.
// Experimental.
type IVpcEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcEndpointRef
type jsiiProxy_IVpcEndpointRef struct {
	internal.Type__constructsIConstruct
}

