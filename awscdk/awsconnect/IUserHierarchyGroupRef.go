package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyGroup.
// Experimental.
type IUserHierarchyGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserHierarchyGroupRef
type jsiiProxy_IUserHierarchyGroupRef struct {
	internal.Type__constructsIConstruct
}

