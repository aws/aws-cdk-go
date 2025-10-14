package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggerDefinition.
// Experimental.
type ILoggerDefinitionRef interface {
	constructs.IConstruct
	// A reference to a LoggerDefinition resource.
	// Experimental.
	LoggerDefinitionRef() *LoggerDefinitionReference
}

// The jsii proxy for ILoggerDefinitionRef
type jsiiProxy_ILoggerDefinitionRef struct {
	internal.Type__constructsIConstruct
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

