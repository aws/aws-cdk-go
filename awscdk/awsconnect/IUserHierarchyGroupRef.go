package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyGroup.
// Experimental.
type IUserHierarchyGroupRef interface {
	constructs.IConstruct
	// A reference to a UserHierarchyGroup resource.
	// Experimental.
	UserHierarchyGroupRef() *UserHierarchyGroupReference
}

// The jsii proxy for IUserHierarchyGroupRef
type jsiiProxy_IUserHierarchyGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserHierarchyGroupRef) UserHierarchyGroupRef() *UserHierarchyGroupReference {
	var returns *UserHierarchyGroupReference
	_jsii_.Get(
		j,
		"userHierarchyGroupRef",
		&returns,
	)
	return returns
}

