package awsapptest


// Specifies an output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &OutputProperty{
//   	File: &OutputFileProperty{
//   		FileLocation: jsii.String("fileLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-output.html
//
type CfnTestCase_OutputProperty struct {
	// The file of the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-output.html#cfn-apptest-testcase-output-file
	//
	File interface{} `field:"required" json:"file" yaml:"file"`
}

