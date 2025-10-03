package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPermissions.
// Experimental.
type ICustomPermissionsRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomPermissionsRef
type jsiiProxy_ICustomPermissionsRef struct {
	internal.Type__constructsIConstruct
}

