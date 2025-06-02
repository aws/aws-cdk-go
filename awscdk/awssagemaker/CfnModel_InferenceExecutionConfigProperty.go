package awssagemaker


// Specifies details about how containers in a multi-container endpoint are run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceExecutionConfigProperty := &InferenceExecutionConfigProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-inferenceexecutionconfig.html
//
type CfnModel_InferenceExecutionConfigProperty struct {
	// How containers in a multi-container are run. The following values are valid.
	//
	// - `Serial` - Containers run as a serial pipeline.
	// - `Direct` - Only the individual container that you specify is run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-inferenceexecutionconfig.html#cfn-sagemaker-model-inferenceexecutionconfig-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

