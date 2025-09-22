package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationRecorder.
// Experimental.
type IConfigurationRecorderRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationRecorder resource.
	// Experimental.
	ConfigurationRecorderRef() *ConfigurationRecorderReference
}

// The jsii proxy for IConfigurationRecorderRef
type jsiiProxy_IConfigurationRecorderRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationRecorderRef) ConfigurationRecorderRef() *ConfigurationRecorderReference {
	var returns *ConfigurationRecorderReference
	_jsii_.Get(
		j,
		"configurationRecorderRef",
		&returns,
	)
	return returns
}

