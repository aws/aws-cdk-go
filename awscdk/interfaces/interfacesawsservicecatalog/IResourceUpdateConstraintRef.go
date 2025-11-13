package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceUpdateConstraint.
// Experimental.
type IResourceUpdateConstraintRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResourceUpdateConstraint resource.
	// Experimental.
	ResourceUpdateConstraintRef() *ResourceUpdateConstraintReference
}

// The jsii proxy for IResourceUpdateConstraintRef
type jsiiProxy_IResourceUpdateConstraintRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IResourceUpdateConstraintRef) ResourceUpdateConstraintRef() *ResourceUpdateConstraintReference {
	var returns *ResourceUpdateConstraintReference
	_jsii_.Get(
		j,
		"resourceUpdateConstraintRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceUpdateConstraintRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceUpdateConstraintRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

