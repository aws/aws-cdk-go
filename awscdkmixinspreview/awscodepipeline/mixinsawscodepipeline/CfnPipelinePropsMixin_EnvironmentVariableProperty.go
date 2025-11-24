package mixinsawscodepipeline


// The environment variables for the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html
//
type CfnPipelinePropsMixin_EnvironmentVariableProperty struct {
	// The environment variable name in the key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html#cfn-codepipeline-pipeline-environmentvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the type of use for the environment variable value.
	//
	// The value can be either `PLAINTEXT` or `SECRETS_MANAGER` . If the value is `SECRETS_MANAGER` , provide the Secrets reference in the EnvironmentVariable value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html#cfn-codepipeline-pipeline-environmentvariable-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The environment variable value in the key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html#cfn-codepipeline-pipeline-environmentvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

