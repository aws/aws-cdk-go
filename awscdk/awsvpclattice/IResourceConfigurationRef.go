package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceConfiguration.
// Experimental.
type IResourceConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceConfigurationRef
type jsiiProxy_IResourceConfigurationRef struct {
	internal.Type__constructsIConstruct
}

