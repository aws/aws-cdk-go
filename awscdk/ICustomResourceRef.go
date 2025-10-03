package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomResource.
// Experimental.
type ICustomResourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomResourceRef
type jsiiProxy_ICustomResourceRef struct {
	internal.Type__constructsIConstruct
}

