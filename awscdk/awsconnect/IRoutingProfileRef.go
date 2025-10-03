package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingProfile.
// Experimental.
type IRoutingProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRoutingProfileRef
type jsiiProxy_IRoutingProfileRef struct {
	internal.Type__constructsIConstruct
}

