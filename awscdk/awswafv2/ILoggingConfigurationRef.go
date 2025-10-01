package awswafv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggingConfiguration.
// Experimental.
type ILoggingConfigurationRef interface {
	constructs.IConstruct
	// A reference to a LoggingConfiguration resource.
	// Experimental.
	LoggingConfigurationRef() *LoggingConfigurationReference
}

// The jsii proxy for ILoggingConfigurationRef
type jsiiProxy_ILoggingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILoggingConfigurationRef) LoggingConfigurationRef() *LoggingConfigurationReference {
	var returns *LoggingConfigurationReference
	_jsii_.Get(
		j,
		"loggingConfigurationRef",
		&returns,
	)
	return returns
}

