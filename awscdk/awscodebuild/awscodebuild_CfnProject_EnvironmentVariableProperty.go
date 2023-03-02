package awscodebuild


// `EnvironmentVariable` is a property of the [AWS CodeBuild Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html) property type that specifies the name and value of an environment variable for an AWS CodeBuild project environment. When you use the environment to run a build, these variables are available for your builds to use. `EnvironmentVariable` contains a list of `EnvironmentVariable` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &environmentVariableProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   }
//
type CfnProject_EnvironmentVariableProperty struct {
	// The name or key of the environment variable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the environment variable.
	//
	// > We strongly discourage the use of `PLAINTEXT` environment variables to store sensitive values, especially AWS secret key IDs and secret access keys. `PLAINTEXT` environment variables can be displayed in plain text using the AWS CodeBuild console and the AWS CLI . For sensitive values, we recommend you use an environment variable of type `PARAMETER_STORE` or `SECRETS_MANAGER` .
	Value *string `field:"required" json:"value" yaml:"value"`
	// The type of environment variable. Valid values include:.
	//
	// - `PARAMETER_STORE` : An environment variable stored in Systems Manager Parameter Store. To learn how to specify a parameter store environment variable, see [env/parameter-store](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.parameter-store) in the *AWS CodeBuild User Guide* .
	// - `PLAINTEXT` : An environment variable in plain text format. This is the default value.
	// - `SECRETS_MANAGER` : An environment variable stored in AWS Secrets Manager . To learn how to specify a secrets manager environment variable, see [env/secrets-manager](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.secrets-manager) in the *AWS CodeBuild User Guide* .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

