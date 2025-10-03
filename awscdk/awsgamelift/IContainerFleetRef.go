package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerFleet.
// Experimental.
type IContainerFleetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContainerFleetRef
type jsiiProxy_IContainerFleetRef struct {
	internal.Type__constructsIConstruct
}

