package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilePermission.
// Experimental.
type IProfilePermissionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProfilePermissionRef
type jsiiProxy_IProfilePermissionRef struct {
	internal.Type__constructsIConstruct
}

