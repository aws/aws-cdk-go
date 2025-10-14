package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggerDefinitionVersion.
// Experimental.
type ILoggerDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a LoggerDefinitionVersion resource.
	// Experimental.
	LoggerDefinitionVersionRef() *LoggerDefinitionVersionReference
}

// The jsii proxy for ILoggerDefinitionVersionRef
type jsiiProxy_ILoggerDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
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

