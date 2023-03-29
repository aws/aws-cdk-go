package awsquicksight


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
type CfnAnalysis_HistogramBinOptionsProperty struct {
	// `CfnAnalysis.HistogramBinOptionsProperty.BinCount`.
	BinCount interface{} `field:"optional" json:"binCount" yaml:"binCount"`
	// `CfnAnalysis.HistogramBinOptionsProperty.BinWidth`.
	BinWidth interface{} `field:"optional" json:"binWidth" yaml:"binWidth"`
	// `CfnAnalysis.HistogramBinOptionsProperty.SelectedBinType`.
	SelectedBinType *string `field:"optional" json:"selectedBinType" yaml:"selectedBinType"`
	// `CfnAnalysis.HistogramBinOptionsProperty.StartValue`.
	StartValue *float64 `field:"optional" json:"startValue" yaml:"startValue"`
}

