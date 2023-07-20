package awstransfer


// Details for a step that invokes an AWS Lambda function.
//
// Consists of the Lambda function's name, target, and timeout (in seconds).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customStepDetailsProperty := &CustomStepDetailsProperty{
//   	Name: jsii.String("name"),
//   	SourceFileLocation: jsii.String("sourceFileLocation"),
//   	Target: jsii.String("target"),
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html
//
type CfnWorkflow_CustomStepDetailsProperty struct {
	// The name of the step, used as an identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file for the workflow.
	//
	// - To use the previous file as the input, enter `${previous.file}` . In this case, this workflow step uses the output file from the previous workflow step as input. This is the default value.
	// - To use the originally uploaded file location as input for this step, enter `${original.file}` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-sourcefilelocation
	//
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
	// The ARN for the Lambda function that is being called.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Timeout, in seconds, for the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-customstepdetails.html#cfn-transfer-workflow-customstepdetails-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

