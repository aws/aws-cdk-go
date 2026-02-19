package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserHierarchyStructure.
// Experimental.
type IUserHierarchyStructureRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserHierarchyStructure resource.
	// Experimental.
	UserHierarchyStructureRef() *UserHierarchyStructureReference
}

// The jsii proxy for IUserHierarchyStructureRef
type jsiiProxy_IUserHierarchyStructureRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IUserHierarchyStructureRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IUserHierarchyStructureRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserHierarchyStructureRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

