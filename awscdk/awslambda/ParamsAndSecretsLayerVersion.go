package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Parameters and Secrets Extension layer version.
//
// Example:
//   import sm "github.com/aws/aws-cdk-go/awscdk"
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   secret := sm.NewSecret(this, jsii.String("Secret"))
//   parameter := ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
//   	ParameterName: jsii.String("mySsmParameterName"),
//   	StringValue: jsii.String("mySsmParameterValue"),
//   })
//
//   paramsAndSecrets := lambda.ParamsAndSecretsLayerVersion_FromVersion(lambda.ParamsAndSecretsVersions_V1_0_103, &ParamsAndSecretsOptions{
//   	CacheSize: jsii.Number(500),
//   	LogLevel: lambda.ParamsAndSecretsLogLevel_DEBUG,
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Architecture: lambda.Architecture_ARM_64(),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	ParamsAndSecrets: ParamsAndSecrets,
//   })
//
//   secret.grantRead(lambdaFunction)
//   parameter.grantRead(lambdaFunction)
//
type ParamsAndSecretsLayerVersion interface {
}

// The jsii proxy struct for ParamsAndSecretsLayerVersion
type jsiiProxy_ParamsAndSecretsLayerVersion struct {
	_ byte // padding
}

// Use a specific version of the Parameters and Secrets Extension to generate a layer version.
func ParamsAndSecretsLayerVersion_FromVersion(version ParamsAndSecretsVersions, options *ParamsAndSecretsOptions) ParamsAndSecretsLayerVersion {
	_init_.Initialize()

	if err := validateParamsAndSecretsLayerVersion_FromVersionParameters(version, options); err != nil {
		panic(err)
	}
	var returns ParamsAndSecretsLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.ParamsAndSecretsLayerVersion",
		"fromVersion",
		[]interface{}{version, options},
		&returns,
	)

	return returns
}

// Use the Parameters and Secrets Extension associated with the provided ARN.
//
// Make sure the ARN is associated
// with the same region and architecture as your function.
// See: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets_lambda.html#retrieving-secrets_lambda_ARNs
//
func ParamsAndSecretsLayerVersion_FromVersionArn(arn *string, options *ParamsAndSecretsOptions) ParamsAndSecretsLayerVersion {
	_init_.Initialize()

	if err := validateParamsAndSecretsLayerVersion_FromVersionArnParameters(arn, options); err != nil {
		panic(err)
	}
	var returns ParamsAndSecretsLayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.ParamsAndSecretsLayerVersion",
		"fromVersionArn",
		[]interface{}{arn, options},
		&returns,
	)

	return returns
}

