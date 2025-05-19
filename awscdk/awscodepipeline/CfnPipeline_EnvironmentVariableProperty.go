package awscodepipeline


// The environment variables for the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html
//
type CfnPipeline_EnvironmentVariableProperty struct {
	// The environment variable name in the key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html#cfn-codepipeline-pipeline-environmentvariable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The environment variable value in the key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-environmentvariable.html#cfn-codepipeline-pipeline-environmentvariable-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

