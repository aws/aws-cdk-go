// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// A secret environment variable.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//
//   secret := secretsmanager.NewSecret(stack, jsii.String("Secret"))
//   parameter := ssm.StringParameter_FromSecureStringParameterAttributes(stack, jsii.String("Parameter"), &SecureStringParameterAttributes{
//   	ParameterName: jsii.String("/name"),
//   	Version: jsii.Number(1),
//   })
//
//   service := apprunner.NewService(stack, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   			EnvironmentSecrets: map[string]secret{
//   				"SECRET": apprunner.*secret_fromSecretsManager(secret),
//   				"PARAMETER": apprunner.*secret_fromSsmParameter(parameter),
//   				"SECRET_ID": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionId": jsii.String("version-id"),
//   				}),
//   				"SECRET_STAGE": apprunner.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   					"versionStage": jsii.String("version-stage"),
//   				}),
//   			},
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.AddSecret(jsii.String("LATER_SECRET"), apprunner.secret_FromSecretsManager(secret, jsii.String("field")))
//
// Experimental.
type Secret interface {
	// The ARN of the secret.
	// Experimental.
	Arn() *string
	// Whether this secret uses a specific JSON field.
	// Experimental.
	HasField() *bool
	// Grants reading the secret to a principal.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for Secret
type jsiiProxy_Secret struct {
	_ byte // padding
}

func (j *jsiiProxy_Secret) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Secret) HasField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasField",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecret_Override(s Secret) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.Secret",
		nil, // no parameters
		s,
	)
}

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
// Experimental.
func Secret_FromSecretsManager(secret awssecretsmanager.ISecret, field *string) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSecretsManagerParameters(secret); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Secret",
		"fromSecretsManager",
		[]interface{}{secret, field},
		&returns,
	)

	return returns
}

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
// Experimental.
func Secret_FromSecretsManagerVersion(secret awssecretsmanager.ISecret, versionInfo *SecretVersionInfo, field *string) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSecretsManagerVersionParameters(secret, versionInfo); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Secret",
		"fromSecretsManagerVersion",
		[]interface{}{secret, versionInfo, field},
		&returns,
	)

	return returns
}

// Creates an environment variable value from a parameter stored in AWS Systems Manager Parameter Store.
// Experimental.
func Secret_FromSsmParameter(parameter awsssm.IParameter) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Secret",
		"fromSsmParameter",
		[]interface{}{parameter},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Secret) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

