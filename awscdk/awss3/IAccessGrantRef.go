package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrant.
// Experimental.
type IAccessGrantRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessGrantRef
type jsiiProxy_IAccessGrantRef struct {
	internal.Type__constructsIConstruct
}

