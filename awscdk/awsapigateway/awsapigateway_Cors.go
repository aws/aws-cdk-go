package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Example:
//   apigateway.NewRestApi(this, jsii.String("api"), &restApiProps{
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: apigateway.cors_ALL_ORIGINS(),
//   		allowMethods: apigateway.*cors_ALL_METHODS(),
//   	},
//   })
//
// Experimental.
type Cors interface {
}

// The jsii proxy struct for Cors
type jsiiProxy_Cors struct {
	_ byte // padding
}

func Cors_ALL_METHODS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.aws_apigateway.Cors",
		"ALL_METHODS",
		&returns,
	)
	return returns
}

func Cors_ALL_ORIGINS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.aws_apigateway.Cors",
		"ALL_ORIGINS",
		&returns,
	)
	return returns
}

func Cors_DEFAULT_HEADERS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.aws_apigateway.Cors",
		"DEFAULT_HEADERS",
		&returns,
	)
	return returns
}

