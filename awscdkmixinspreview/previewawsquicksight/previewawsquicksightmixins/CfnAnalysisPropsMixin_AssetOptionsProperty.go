package previewawsquicksightmixins


// An array of analysis level configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetOptionsProperty := &AssetOptionsProperty{
//   	Timezone: jsii.String("timezone"),
//   	WeekStart: jsii.String("weekStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html
//
type CfnAnalysisPropsMixin_AssetOptionsProperty struct {
	// Determines the timezone for the analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Determines the week start day for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-assetoptions.html#cfn-quicksight-analysis-assetoptions-weekstart
	//
	WeekStart *string `field:"optional" json:"weekStart" yaml:"weekStart"`
}

