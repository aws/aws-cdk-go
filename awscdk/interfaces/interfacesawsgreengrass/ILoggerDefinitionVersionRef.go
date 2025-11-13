package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggerDefinitionVersion.
// Experimental.
type ILoggerDefinitionVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LoggerDefinitionVersion resource.
	// Experimental.
	LoggerDefinitionVersionRef() *LoggerDefinitionVersionReference
}

// The jsii proxy for ILoggerDefinitionVersionRef
type jsiiProxy_ILoggerDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILoggerDefinitionVersionRef) LoggerDefinitionVersionRef() *LoggerDefinitionVersionReference {
	var returns *LoggerDefinitionVersionReference
	_jsii_.Get(
		j,
		"loggerDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggerDefinitionVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggerDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

