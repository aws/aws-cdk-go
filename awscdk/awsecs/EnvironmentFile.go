package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Constructs for types of environment files.
//
// Example:
//   var secret secret
//   var dbSecret secret
//   var parameter stringParameter
//   var taskDefinition taskDefinition
//   var s3Bucket bucket
//
//
//   newContainer := taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(1024),
//   	Environment: map[string]*string{
//   		 // clear text, not for sensitive data
//   		"STAGE": jsii.String("prod"),
//   	},
//   	EnvironmentFiles: []environmentFile{
//   		ecs.*environmentFile_FromAsset(jsii.String("./demo-env-file.env")),
//   		ecs.*environmentFile_FromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
//   	},
//   	Secrets: map[string]secret{
//   		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
//   		"SECRET": ecs.*secret_fromSecretsManager(secret),
//   		"DB_PASSWORD": ecs.*secret_fromSecretsManager(dbSecret, jsii.String("password")),
//   		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
//   		"API_KEY": ecs.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
//   			"versionId": jsii.String("12345"),
//   		}, jsii.String("apiKey")),
//   		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
//   		"PARAMETER": ecs.*secret_fromSsmParameter(parameter),
//   	},
//   })
//   newContainer.AddEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
//   newContainer.AddSecret(jsii.String("API_KEY"), ecs.secret_FromSecretsManager(secret))
//   newContainer.AddSecret(jsii.String("DB_PASSWORD"), ecs.secret_FromSecretsManager(secret, jsii.String("password")))
//
type EnvironmentFile interface {
	// Called when the container is initialized to allow this object to bind to the stack.
	Bind(scope constructs.Construct) *EnvironmentFileConfig
}

// The jsii proxy struct for EnvironmentFile
type jsiiProxy_EnvironmentFile struct {
	_ byte // padding
}

func NewEnvironmentFile_Override(e EnvironmentFile) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		nil, // no parameters
		e,
	)
}

// Loads the environment file from a local disk path.
func EnvironmentFile_FromAsset(path *string, options *awss3assets.AssetOptions) AssetEnvironmentFile {
	_init_.Initialize()

	if err := validateEnvironmentFile_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetEnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Loads the environment file from an S3 bucket.
//
// Returns: `S3EnvironmentFile` associated with the specified S3 object.
func EnvironmentFile_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3EnvironmentFile {
	_init_.Initialize()

	if err := validateEnvironmentFile_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3EnvironmentFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EnvironmentFile",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentFile) Bind(scope constructs.Construct) *EnvironmentFileConfig {
	if err := e.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *EnvironmentFileConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

