package awsacmpca

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Permission.
// Experimental.
type IPermissionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPermissionRef
type jsiiProxy_IPermissionRef struct {
	internal.Type__constructsIConstruct
}

