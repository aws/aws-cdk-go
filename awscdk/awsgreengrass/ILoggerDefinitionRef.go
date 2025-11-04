package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggerDefinition.
// Experimental.
type ILoggerDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LoggerDefinition resource.
	// Experimental.
	LoggerDefinitionRef() *LoggerDefinitionReference
}

// The jsii proxy for ILoggerDefinitionRef
type jsiiProxy_ILoggerDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILoggerDefinitionRef) LoggerDefinitionRef() *LoggerDefinitionReference {
	var returns *LoggerDefinitionReference
	_jsii_.Get(
		j,
		"loggerDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggerDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggerDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

