package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Defines an AWS Lambda validator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   lambdaValidator := appconfig_alpha.LambdaValidator_FromFunction(function_)
//
// Deprecated.
type LambdaValidator interface {
	IValidator
	// The content of the validator.
	// Deprecated.
	Content() *string
	// The type of validator.
	// Deprecated.
	Type() ValidatorType
}

// The jsii proxy struct for LambdaValidator
type jsiiProxy_LambdaValidator struct {
	jsiiProxy_IValidator
}

func (j *jsiiProxy_LambdaValidator) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaValidator) Type() ValidatorType {
	var returns ValidatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Deprecated.
func NewLambdaValidator_Override(l LambdaValidator) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.LambdaValidator",
		nil, // no parameters
		l,
	)
}

// Defines an AWS Lambda validator from a Lambda function.
//
// This will call
// `addPermission` to your function to grant AWS AppConfig permissions.
// Deprecated.
func LambdaValidator_FromFunction(func_ awslambda.Function) LambdaValidator {
	_init_.Initialize()

	if err := validateLambdaValidator_FromFunctionParameters(func_); err != nil {
		panic(err)
	}
	var returns LambdaValidator

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.LambdaValidator",
		"fromFunction",
		[]interface{}{func_},
		&returns,
	)

	return returns
}

