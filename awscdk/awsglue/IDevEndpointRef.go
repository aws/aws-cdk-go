package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DevEndpoint.
// Experimental.
type IDevEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDevEndpointRef
type jsiiProxy_IDevEndpointRef struct {
	internal.Type__constructsIConstruct
}

