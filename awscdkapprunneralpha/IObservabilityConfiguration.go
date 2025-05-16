package awscdkapprunneralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
)

// Represents the App Runner Observability configuration.
// Experimental.
type IObservabilityConfiguration interface {
	awscdk.IResource
	// The ARN of the Observability configuration.
	// Experimental.
	ObservabilityConfigurationArn() *string
	// The Name of the Observability configuration.
	// Experimental.
	ObservabilityConfigurationName() *string
	// The revision of the Observability configuration.
	// Experimental.
	ObservabilityConfigurationRevision() *float64
}

// The jsii proxy for IObservabilityConfiguration
type jsiiProxy_IObservabilityConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IObservabilityConfiguration) ObservabilityConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObservabilityConfiguration) ObservabilityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"observabilityConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObservabilityConfiguration) ObservabilityConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"observabilityConfigurationRevision",
		&returns,
	)
	return returns
}

