package awsimagebuilder


// The condition that the output AMI's scan results must meet for the image to pass.
//
// When you set it, the test phase waits for the scan to finish, and Image Builder fails the image if a counted finding exceeds MaximumSeverityLabel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scanResultConditionProperty := &ScanResultConditionProperty{
//   	MaximumSeverityLabel: jsii.String("maximumSeverityLabel"),
//
//   	// the properties below are optional
//   	FixAvailable: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-scanresultcondition.html
//
type CfnImage_ScanResultConditionProperty struct {
	// The maximum severity allowed for findings on the output AMI.
	//
	// A finding above this severity fails the image, unless FixAvailable excludes it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-scanresultcondition.html#cfn-imagebuilder-image-scanresultcondition-maximumseveritylabel
	//
	MaximumSeverityLabel *string `field:"required" json:"maximumSeverityLabel" yaml:"maximumSeverityLabel"`
	// Specifies whether only findings that have a fix available count toward MaximumSeverityLabel.
	//
	// When false, Image Builder counts every finding. The default is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-image-scanresultcondition.html#cfn-imagebuilder-image-scanresultcondition-fixavailable
	//
	FixAvailable interface{} `field:"optional" json:"fixAvailable" yaml:"fixAvailable"`
}

