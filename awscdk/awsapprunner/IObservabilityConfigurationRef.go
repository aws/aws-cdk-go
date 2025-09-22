package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ObservabilityConfiguration.
// Experimental.
type IObservabilityConfigurationRef interface {
	constructs.IConstruct
	// A reference to a ObservabilityConfiguration resource.
	// Experimental.
	ObservabilityConfigurationRef() *ObservabilityConfigurationReference
}

// The jsii proxy for IObservabilityConfigurationRef
type jsiiProxy_IObservabilityConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IObservabilityConfigurationRef) ObservabilityConfigurationRef() *ObservabilityConfigurationReference {
	var returns *ObservabilityConfigurationReference
	_jsii_.Get(
		j,
		"observabilityConfigurationRef",
		&returns,
	)
	return returns
}

