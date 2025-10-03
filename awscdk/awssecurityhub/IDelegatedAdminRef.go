package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DelegatedAdmin.
// Experimental.
type IDelegatedAdminRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDelegatedAdminRef
type jsiiProxy_IDelegatedAdminRef struct {
	internal.Type__constructsIConstruct
}

