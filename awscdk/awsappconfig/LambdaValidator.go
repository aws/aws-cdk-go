package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Defines an AWS Lambda validator.
//
// Example:
//   var application application
//   var fn function
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Validators: []iValidator{
//   		appconfig.JsonSchemaValidator_FromFile(jsii.String("schema.json")),
//   		appconfig.LambdaValidator_FromFunction(fn),
//   	},
//   })
//
type LambdaValidator interface {
	IValidator
	// The content of the validator.
	Content() *string
	// The type of validator.
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


func NewLambdaValidator_Override(l LambdaValidator) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.LambdaValidator",
		nil, // no parameters
		l,
	)
}

// Defines an AWS Lambda validator from a Lambda function.
//
// This will call
// `addPermission` to your function to grant AWS AppConfig permissions.
func LambdaValidator_FromFunction(func_ awslambda.Function) LambdaValidator {
	_init_.Initialize()

	if err := validateLambdaValidator_FromFunctionParameters(func_); err != nil {
		panic(err)
	}
	var returns LambdaValidator

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appconfig.LambdaValidator",
		"fromFunction",
		[]interface{}{func_},
		&returns,
	)

	return returns
}

