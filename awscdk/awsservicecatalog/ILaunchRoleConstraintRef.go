package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchRoleConstraint.
// Experimental.
type ILaunchRoleConstraintRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILaunchRoleConstraintRef
type jsiiProxy_ILaunchRoleConstraintRef struct {
	internal.Type__constructsIConstruct
}

