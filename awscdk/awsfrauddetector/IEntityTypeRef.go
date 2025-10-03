package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EntityType.
// Experimental.
type IEntityTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEntityTypeRef
type jsiiProxy_IEntityTypeRef struct {
	internal.Type__constructsIConstruct
}

