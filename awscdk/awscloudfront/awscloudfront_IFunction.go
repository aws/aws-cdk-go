package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
)

// Represents a CloudFront Function.
type IFunction interface {
	awscdk.IResource
	// The ARN of the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
}

// The jsii proxy for IFunction
type jsiiProxy_IFunction struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

