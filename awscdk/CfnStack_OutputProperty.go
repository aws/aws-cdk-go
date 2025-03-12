package awscdk


// The Output data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &OutputProperty{
//   	Description: jsii.String("description"),
//   	ExportName: jsii.String("exportName"),
//   	OutputKey: jsii.String("outputKey"),
//   	OutputValue: jsii.String("outputValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stack-output.html
//
type CfnStack_OutputProperty struct {
	// User defined description associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stack-output.html#cfn-cloudformation-stack-output-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the export associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stack-output.html#cfn-cloudformation-stack-output-exportname
	//
	ExportName *string `field:"optional" json:"exportName" yaml:"exportName"`
	// The key associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stack-output.html#cfn-cloudformation-stack-output-outputkey
	//
	OutputKey *string `field:"optional" json:"outputKey" yaml:"outputKey"`
	// The value associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stack-output.html#cfn-cloudformation-stack-output-outputvalue
	//
	OutputValue *string `field:"optional" json:"outputValue" yaml:"outputValue"`
}

