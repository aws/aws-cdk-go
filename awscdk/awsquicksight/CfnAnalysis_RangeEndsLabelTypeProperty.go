package awsquicksight


// The range ends label type of a data path label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangeEndsLabelTypeProperty := &RangeEndsLabelTypeProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-rangeendslabeltype.html
//
type CfnAnalysis_RangeEndsLabelTypeProperty struct {
	// The visibility of the range ends label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-rangeendslabeltype.html#cfn-quicksight-analysis-rangeendslabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

