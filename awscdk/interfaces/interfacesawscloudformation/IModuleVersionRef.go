package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleVersion.
// Experimental.
type IModuleVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ModuleVersion resource.
	// Experimental.
	ModuleVersionRef() *ModuleVersionReference
}

// The jsii proxy for IModuleVersionRef
type jsiiProxy_IModuleVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IModuleVersionRef) ModuleVersionRef() *ModuleVersionReference {
	var returns *ModuleVersionReference
	_jsii_.Get(
		j,
		"moduleVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModuleVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModuleVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

