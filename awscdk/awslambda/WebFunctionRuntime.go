package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Runtime environment for WebFunction.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("PerRegionApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_PerRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		}),
//   	},
//   })
//
type WebFunctionRuntime interface {
	// The name of this runtime, as expected by the WebFunction resource.
	Name() *string
	// Returns the runtime name.
	ToString() *string
}

// The jsii proxy struct for WebFunctionRuntime
type jsiiProxy_WebFunctionRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_WebFunctionRuntime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Use a custom runtime string not yet supported by the CDK.
//
// This allows customers to use a newly released runtime without waiting
// for a CDK library update.
func WebFunctionRuntime_Of(name *string) WebFunctionRuntime {
	_init_.Initialize()

	if err := validateWebFunctionRuntime_OfParameters(name); err != nil {
		panic(err)
	}
	var returns WebFunctionRuntime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.WebFunctionRuntime",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func WebFunctionRuntime_NODEJS_24_X() WebFunctionRuntime {
	_init_.Initialize()
	var returns WebFunctionRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.WebFunctionRuntime",
		"NODEJS_24_X",
		&returns,
	)
	return returns
}

func WebFunctionRuntime_NODEJS_LATEST() WebFunctionRuntime {
	_init_.Initialize()
	var returns WebFunctionRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.WebFunctionRuntime",
		"NODEJS_LATEST",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WebFunctionRuntime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

