package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// An additional HTTP parameter to send along with the OAuth request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//
//   httpParameter := awscdk.Aws_events.HttpParameter_FromSecret(secretValue)
//
type HttpParameter interface {
}

// The jsii proxy struct for HttpParameter
type jsiiProxy_HttpParameter struct {
	_ byte // padding
}

func NewHttpParameter_Override(h HttpParameter) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.HttpParameter",
		nil, // no parameters
		h,
	)
}

// Make an OAuthParameter from a secret.
func HttpParameter_FromSecret(value awscdk.SecretValue) HttpParameter {
	_init_.Initialize()

	if err := validateHttpParameter_FromSecretParameters(value); err != nil {
		panic(err)
	}
	var returns HttpParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.HttpParameter",
		"fromSecret",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Make an OAuthParameter from a string value.
//
// The value is not treated as a secret.
func HttpParameter_FromString(value *string) HttpParameter {
	_init_.Initialize()

	if err := validateHttpParameter_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns HttpParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.HttpParameter",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

