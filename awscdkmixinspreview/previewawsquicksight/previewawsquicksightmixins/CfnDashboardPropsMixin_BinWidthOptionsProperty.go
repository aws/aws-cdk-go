package previewawsquicksightmixins


// The options that determine the bin width of a histogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   binWidthOptionsProperty := &BinWidthOptionsProperty{
//   	BinCountLimit: jsii.Number(123),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-binwidthoptions.html
//
type CfnDashboardPropsMixin_BinWidthOptionsProperty struct {
	// The options that determine the bin count limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-binwidthoptions.html#cfn-quicksight-dashboard-binwidthoptions-bincountlimit
	//
	BinCountLimit *float64 `field:"optional" json:"binCountLimit" yaml:"binCountLimit"`
	// The options that determine the bin width value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-binwidthoptions.html#cfn-quicksight-dashboard-binwidthoptions-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

