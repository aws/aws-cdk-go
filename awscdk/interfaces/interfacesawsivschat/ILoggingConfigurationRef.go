package interfacesawsivschat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsivschat/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggingConfiguration.
// Experimental.
type ILoggingConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LoggingConfiguration resource.
	// Experimental.
	LoggingConfigurationRef() *LoggingConfigurationReference
}

// The jsii proxy for ILoggingConfigurationRef
type jsiiProxy_ILoggingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ILoggingConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoggingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

