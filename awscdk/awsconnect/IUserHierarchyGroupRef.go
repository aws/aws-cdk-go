package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyGroup.
// Experimental.
type IUserHierarchyGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserHierarchyGroup resource.
	// Experimental.
	UserHierarchyGroupRef() *UserHierarchyGroupReference
}

// The jsii proxy for IUserHierarchyGroupRef
type jsiiProxy_IUserHierarchyGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IUserHierarchyGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserHierarchyGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

