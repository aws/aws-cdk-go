// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
)

// Properties of a compute environment.
// Experimental.
type IComputeEnvironment interface {
	awscdk.IResource
	// The ARN of this compute environment.
	// Experimental.
	ComputeEnvironmentArn() *string
	// The name of this compute environment.
	// Experimental.
	ComputeEnvironmentName() *string
}

// The jsii proxy for IComputeEnvironment
type jsiiProxy_IComputeEnvironment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

