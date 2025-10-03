package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceLinkedRole.
// Experimental.
type IServiceLinkedRoleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceLinkedRoleRef
type jsiiProxy_IServiceLinkedRoleRef struct {
	internal.Type__constructsIConstruct
}

