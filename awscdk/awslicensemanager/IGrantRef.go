package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Grant.
// Experimental.
type IGrantRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGrantRef
type jsiiProxy_IGrantRef struct {
	internal.Type__constructsIConstruct
}

