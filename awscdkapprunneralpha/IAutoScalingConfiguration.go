package awscdkapprunneralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
)

// Represents the App Runner Auto Scaling Configuration.
// Experimental.
type IAutoScalingConfiguration interface {
	awscdk.IResource
	// The ARN of the Auto Scaling Configuration.
	// Experimental.
	AutoScalingConfigurationArn() *string
	// The Name of the Auto Scaling Configuration.
	// Experimental.
	AutoScalingConfigurationName() *string
	// The revision of the Auto Scaling Configuration.
	// Experimental.
	AutoScalingConfigurationRevision() *float64
}

// The jsii proxy for IAutoScalingConfiguration
type jsiiProxy_IAutoScalingConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAutoScalingConfiguration) AutoScalingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingConfiguration) AutoScalingConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingConfiguration) AutoScalingConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingConfigurationRevision",
		&returns,
	)
	return returns
}

