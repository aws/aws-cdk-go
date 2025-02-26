package awsquicksight


// The options that determine the thousands separator configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thousandSeparatorOptionsProperty := &ThousandSeparatorOptionsProperty{
//   	Symbol: jsii.String("symbol"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-thousandseparatoroptions.html
//
type CfnAnalysis_ThousandSeparatorOptionsProperty struct {
	// Determines the thousands separator symbol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-thousandseparatoroptions.html#cfn-quicksight-analysis-thousandseparatoroptions-symbol
	//
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
	// Determines the visibility of the thousands separator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-thousandseparatoroptions.html#cfn-quicksight-analysis-thousandseparatoroptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

