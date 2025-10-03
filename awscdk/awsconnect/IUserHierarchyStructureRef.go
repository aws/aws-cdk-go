package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyStructure.
// Experimental.
type IUserHierarchyStructureRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserHierarchyStructureRef
type jsiiProxy_IUserHierarchyStructureRef struct {
	internal.Type__constructsIConstruct
}

