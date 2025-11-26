package previewawsgreengrassmixins


// The default configuration that applies to all Lambda functions in the function definition version.
//
// Individual Lambda functions can override these settings.
//
// In an CloudFormation template, `DefaultConfig` is a property of the [`FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultConfigProperty := &DefaultConfigProperty{
//   	Execution: &ExecutionProperty{
//   		IsolationMode: jsii.String("isolationMode"),
//   		RunAs: &RunAsProperty{
//   			Gid: jsii.Number(123),
//   			Uid: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-defaultconfig.html
//
type CfnFunctionDefinitionPropsMixin_DefaultConfigProperty struct {
	// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-defaultconfig.html#cfn-greengrass-functiondefinition-defaultconfig-execution
	//
	Execution interface{} `field:"optional" json:"execution" yaml:"execution"`
}

