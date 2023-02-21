package awssagemaker


// Specifies details about how containers in a multi-container endpoint are run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceExecutionConfigProperty := &inferenceExecutionConfigProperty{
//   	mode: jsii.String("mode"),
//   }
//
type CfnModel_InferenceExecutionConfigProperty struct {
	// How containers in a multi-container are run. The following values are valid.
	//
	// - `Serial` - Containers run as a serial pipeline.
	// - `Direct` - Only the individual container that you specify is run.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

