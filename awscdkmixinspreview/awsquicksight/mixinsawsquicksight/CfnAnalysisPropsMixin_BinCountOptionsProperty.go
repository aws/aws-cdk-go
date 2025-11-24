package mixinsawsquicksight


// The options that determine the bin count of a histogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   binCountOptionsProperty := &BinCountOptionsProperty{
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bincountoptions.html
//
type CfnAnalysisPropsMixin_BinCountOptionsProperty struct {
	// The options that determine the bin count value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bincountoptions.html#cfn-quicksight-analysis-bincountoptions-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

