package awsquicksight


// The options that determine the presentation of the secondary value of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryValueOptionsProperty := &SecondaryValueOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-secondaryvalueoptions.html
//
type CfnAnalysis_SecondaryValueOptionsProperty struct {
	// Determines the visibility of the secondary value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-secondaryvalueoptions.html#cfn-quicksight-analysis-secondaryvalueoptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

