package awsquicksight


// The options that determine the presentation of histogram bins.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   histogramBinOptionsProperty := &HistogramBinOptionsProperty{
//   	BinCount: &BinCountOptionsProperty{
//   		Value: jsii.Number(123),
//   	},
//   	BinWidth: &BinWidthOptionsProperty{
//   		BinCountLimit: jsii.Number(123),
//   		Value: jsii.Number(123),
//   	},
//   	SelectedBinType: jsii.String("selectedBinType"),
//   	StartValue: jsii.Number(123),
//   }
//
type CfnTemplate_HistogramBinOptionsProperty struct {
	// The options that determine the bin count of a histogram.
	BinCount interface{} `field:"optional" json:"binCount" yaml:"binCount"`
	// The options that determine the bin width of a histogram.
	BinWidth interface{} `field:"optional" json:"binWidth" yaml:"binWidth"`
	// The options that determine the selected bin type.
	SelectedBinType *string `field:"optional" json:"selectedBinType" yaml:"selectedBinType"`
	// The options that determine the bin start value.
	StartValue *float64 `field:"optional" json:"startValue" yaml:"startValue"`
}

