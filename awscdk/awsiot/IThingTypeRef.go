package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingType.
// Experimental.
type IThingTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThingTypeRef
type jsiiProxy_IThingTypeRef struct {
	internal.Type__constructsIConstruct
}

