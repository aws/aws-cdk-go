package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleDefaultVersion.
// Experimental.
type IModuleDefaultVersionRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a ModuleDefaultVersion resource.
	// Experimental.
	ModuleDefaultVersionRef() *ModuleDefaultVersionReference
}

// The jsii proxy for IModuleDefaultVersionRef
type jsiiProxy_IModuleDefaultVersionRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
}

func (j *jsiiProxy_IModuleDefaultVersionRef) ModuleDefaultVersionRef() *ModuleDefaultVersionReference {
	var returns *ModuleDefaultVersionReference
	_jsii_.Get(
		j,
		"moduleDefaultVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModuleDefaultVersionRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModuleDefaultVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

