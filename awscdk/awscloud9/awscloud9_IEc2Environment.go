package awscloud9

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloud9/internal"
)

// A Cloud9 Environment.
// Experimental.
type IEc2Environment interface {
	awscdk.IResource
	// The arn of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The name of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentName() *string
}

// The jsii proxy for IEc2Environment
type jsiiProxy_IEc2Environment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

