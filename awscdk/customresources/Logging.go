package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class used to configure Logging during AwsCustomResource SDK calls.
//
// Example:
//   getParameter := cr.NewAwsCustomResource(this, jsii.String("GetParameter"), &AwsCustomResourceProps{
//   	OnUpdate: &AwsSdkCall{
//   		Service: jsii.String("SSM"),
//   		Action: jsii.String("GetParameter"),
//   		Parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(date.now().toString()),
//   		Logging: cr.Logging_WithDataHidden(),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type Logging interface {
}

// The jsii proxy struct for Logging
type jsiiProxy_Logging struct {
	_ byte // padding
}

func NewLogging_Override(l Logging, props *LoggingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.Logging",
		[]interface{}{props},
		l,
	)
}

// Enables logging of all logged data in the lambda handler.
//
// This includes the event object, the API call response, all fields in the response object
// returned by the lambda, and any errors encountered.
func Logging_All() Logging {
	_init_.Initialize()

	var returns Logging

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.Logging",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Hides logging of data associated with the API call response.
//
// This includes hiding the raw API
// call response and the `Data` field associated with the lambda handler response.
func Logging_WithDataHidden() Logging {
	_init_.Initialize()

	var returns Logging

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.Logging",
		"withDataHidden",
		nil, // no parameters
		&returns,
	)

	return returns
}

