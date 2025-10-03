package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Thing.
// Experimental.
type IThingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThingRef
type jsiiProxy_IThingRef struct {
	internal.Type__constructsIConstruct
}

