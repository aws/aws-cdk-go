package awslakeformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permissions.
// Experimental.
type IPermissionsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPermissionsRef
type jsiiProxy_IPermissionsRef struct {
	internal.Type__constructsIConstruct
}

