package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

// A secret environment variable.
//
// Example:
//   var secret secret
//   var parameter stringParameter
//
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_Firelens(&FireLensLogDriverProps{
//   		Options: map[string]interface{}{
//   		},
//   		SecretOptions: map[string]secret{
//   			 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store
//   			"apikey": ecs.*secret_fromSecretsManager(secret),
//   			"host": ecs.*secret_fromSsmParameter(parameter),
//   		},
//   	}),
//   })
//
type Secret interface {
	// The ARN of the secret.
	Arn() *string
	// Whether this secret uses a specific JSON field.
	HasField() *bool
	// Grants reading the secret to a principal.
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


func NewSecret_Override(s Secret) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Secret",
		nil, // no parameters
		s,
	)
}

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
func Secret_FromSecretsManager(secret awssecretsmanager.ISecret, field *string) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSecretsManagerParameters(secret); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
		"fromSecretsManager",
		[]interface{}{secret, field},
		&returns,
	)

	return returns
}

// Creates a environment variable value from a secret stored in AWS Secrets Manager.
func Secret_FromSecretsManagerVersion(secret awssecretsmanager.ISecret, versionInfo *SecretVersionInfo, field *string) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSecretsManagerVersionParameters(secret, versionInfo); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
		"fromSecretsManagerVersion",
		[]interface{}{secret, versionInfo, field},
		&returns,
	)

	return returns
}

// Creates an environment variable value from a parameter stored in AWS Systems Manager Parameter Store.
func Secret_FromSsmParameter(parameter awsssm.IParameter) Secret {
	_init_.Initialize()

	if err := validateSecret_FromSsmParameterParameters(parameter); err != nil {
		panic(err)
	}
	var returns Secret

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Secret",
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

