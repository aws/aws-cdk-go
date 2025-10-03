package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginEndpoint.
// Experimental.
type IOriginEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IOriginEndpointRef
type jsiiProxy_IOriginEndpointRef struct {
	internal.Type__constructsIConstruct
}

