package awsamplifyuibuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Component.
// Experimental.
type IComponentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComponentRef
type jsiiProxy_IComponentRef struct {
	internal.Type__constructsIConstruct
}

