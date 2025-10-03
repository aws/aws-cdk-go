package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingGroup.
// Experimental.
type IThingGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThingGroupRef
type jsiiProxy_IThingGroupRef struct {
	internal.Type__constructsIConstruct
}

