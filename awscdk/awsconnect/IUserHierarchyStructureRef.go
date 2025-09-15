package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyStructure.
// Experimental.
type IUserHierarchyStructureRef interface {
	constructs.IConstruct
	// A reference to a UserHierarchyStructure resource.
	// Experimental.
	UserHierarchyStructureRef() *UserHierarchyStructureReference
}

// The jsii proxy for IUserHierarchyStructureRef
type jsiiProxy_IUserHierarchyStructureRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserHierarchyStructureRef) UserHierarchyStructureRef() *UserHierarchyStructureReference {
	var returns *UserHierarchyStructureReference
	_jsii_.Get(
		j,
		"userHierarchyStructureRef",
		&returns,
	)
	return returns
}

