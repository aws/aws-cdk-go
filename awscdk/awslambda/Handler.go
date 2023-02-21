package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Lambda function handler.
type Handler interface {
}

// The jsii proxy struct for Handler
type jsiiProxy_Handler struct {
	_ byte // padding
}

func Handler_FROM_IMAGE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Handler",
		"FROM_IMAGE",
		&returns,
	)
	return returns
}

