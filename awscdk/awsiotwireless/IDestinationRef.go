package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Destination.
// Experimental.
type IDestinationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDestinationRef
type jsiiProxy_IDestinationRef struct {
	internal.Type__constructsIConstruct
}

