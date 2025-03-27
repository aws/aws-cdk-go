package awscodebuild


// `EnvironmentVariable` is a property of the [AWS CodeBuild Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html) property type that specifies the name and value of an environment variable for an AWS CodeBuild project environment. When you use the environment to run a build, these variables are available for your builds to use. `EnvironmentVariable` contains a list of `EnvironmentVariable` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
//
type CfnProject_EnvironmentVariableProperty struct {
	// The name or key of the environment variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the environment variable.
	//
	// > We strongly discourage the use of `PLAINTEXT` environment variables to store sensitive values, especially AWS secret key IDs. `PLAINTEXT` environment variables can be displayed in plain text using the AWS CodeBuild console and the AWS CLI . For sensitive values, we recommend you use an environment variable of type `PARAMETER_STORE` or `SECRETS_MANAGER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
	// The type of environment variable. Valid values include:.
	//
	// - `PARAMETER_STORE` : An environment variable stored in Systems Manager Parameter Store. For environment variables of this type, specify the name of the parameter as the `value` of the EnvironmentVariable. The parameter value will be substituted for the name at runtime. You can also define Parameter Store environment variables in the buildspec. To learn how to do so, see [env/parameter-store](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.parameter-store) in the *AWS CodeBuild User Guide* .
	// - `PLAINTEXT` : An environment variable in plain text format. This is the default value.
	// - `SECRETS_MANAGER` : An environment variable stored in AWS Secrets Manager . For environment variables of this type, specify the name of the secret as the `value` of the EnvironmentVariable. The secret value will be substituted for the name at runtime. You can also define AWS Secrets Manager environment variables in the buildspec. To learn how to do so, see [env/secrets-manager](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.secrets-manager) in the *AWS CodeBuild User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

