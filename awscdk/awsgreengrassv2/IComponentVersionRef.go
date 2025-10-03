package awsgreengrassv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrassv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentVersion.
// Experimental.
type IComponentVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComponentVersionRef
type jsiiProxy_IComponentVersionRef struct {
	internal.Type__constructsIConstruct
}

