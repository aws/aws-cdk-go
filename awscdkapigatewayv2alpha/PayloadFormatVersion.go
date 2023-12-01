package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Payload format version for lambda proxy integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   payloadFormatVersion := apigatewayv2_alpha.PayloadFormatVersion_Custom(jsii.String("version"))
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
//
// Deprecated.
type PayloadFormatVersion interface {
	// version as a string.
	// Deprecated.
	Version() *string
}

// The jsii proxy struct for PayloadFormatVersion
type jsiiProxy_PayloadFormatVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_PayloadFormatVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// A custom payload version.
//
// Typically used if there is a version number that the CDK doesn't support yet.
// Deprecated.
func PayloadFormatVersion_Custom(version *string) PayloadFormatVersion {
	_init_.Initialize()

	if err := validatePayloadFormatVersion_CustomParameters(version); err != nil {
		panic(err)
	}
	var returns PayloadFormatVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"custom",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func PayloadFormatVersion_VERSION_1_0() PayloadFormatVersion {
	_init_.Initialize()
	var returns PayloadFormatVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"VERSION_1_0",
		&returns,
	)
	return returns
}

func PayloadFormatVersion_VERSION_2_0() PayloadFormatVersion {
	_init_.Initialize()
	var returns PayloadFormatVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-apigatewayv2-alpha.PayloadFormatVersion",
		"VERSION_2_0",
		&returns,
	)
	return returns
}

