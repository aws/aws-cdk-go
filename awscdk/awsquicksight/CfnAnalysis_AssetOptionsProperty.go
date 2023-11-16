package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetOptionsProperty := &AssetOptionsProperty{
//   	Timezone: jsii.String("timezone"),
//   	WeekStart: jsii.String("weekStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html
//
type CfnAnalysis_AssetOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-weekstart
	//
	WeekStart *string `field:"optional" json:"weekStart" yaml:"weekStart"`
}

