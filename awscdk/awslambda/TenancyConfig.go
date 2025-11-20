package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specify the tenant isolation mode for Lambda functions.
//
// This is incompatible with:
// - SnapStart
// - Provisioned Concurrency
// - Function URLs
// - Most Event sources (only API Gateway is supported).
//
// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	TenancyConfig: lambda.TenancyConfig_PER_TENANT(),
//   })
//
type TenancyConfig interface {
	// The CloudFormation property for tenancy configuration.
	TenancyConfigProperty() *CfnFunction_TenancyConfigProperty
}

// The jsii proxy struct for TenancyConfig
type jsiiProxy_TenancyConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_TenancyConfig) TenancyConfigProperty() *CfnFunction_TenancyConfigProperty {
	var returns *CfnFunction_TenancyConfigProperty
	_jsii_.Get(
		j,
		"tenancyConfigProperty",
		&returns,
	)
	return returns
}


func NewTenancyConfig(mode *string) TenancyConfig {
	_init_.Initialize()

	if err := validateNewTenancyConfigParameters(mode); err != nil {
		panic(err)
	}
	j := jsiiProxy_TenancyConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.TenancyConfig",
		[]interface{}{mode},
		&j,
	)

	return &j
}

func NewTenancyConfig_Override(t TenancyConfig, mode *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.TenancyConfig",
		[]interface{}{mode},
		t,
	)
}

func TenancyConfig_PER_TENANT() TenancyConfig {
	_init_.Initialize()
	var returns TenancyConfig
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.TenancyConfig",
		"PER_TENANT",
		&returns,
	)
	return returns
}

