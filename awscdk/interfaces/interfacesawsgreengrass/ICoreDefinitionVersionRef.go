package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreDefinitionVersion.
// Experimental.
type ICoreDefinitionVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CoreDefinitionVersion resource.
	// Experimental.
	CoreDefinitionVersionRef() *CoreDefinitionVersionReference
}

// The jsii proxy for ICoreDefinitionVersionRef
type jsiiProxy_ICoreDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICoreDefinitionVersionRef) CoreDefinitionVersionRef() *CoreDefinitionVersionReference {
	var returns *CoreDefinitionVersionReference
	_jsii_.Get(
		j,
		"coreDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICoreDefinitionVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICoreDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

