package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
)

// A secret environment variable.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	memoryLimitMiB: jsii.Number(1024),
//   	environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	environmentFiles: []environmentFile{
//   		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret.fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
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
		"monocdk.aws_ecs.Secret",
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
		"monocdk.aws_ecs.Secret",
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
		"monocdk.aws_ecs.Secret",
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
		"monocdk.aws_ecs.Secret",
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

