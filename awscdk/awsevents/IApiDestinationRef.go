package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiDestination.
// Experimental.
type IApiDestinationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApiDestinationRef
type jsiiProxy_IApiDestinationRef struct {
	internal.Type__constructsIConstruct
}

