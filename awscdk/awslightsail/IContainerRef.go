package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Container.
// Experimental.
type IContainerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContainerRef
type jsiiProxy_IContainerRef struct {
	internal.Type__constructsIConstruct
}

