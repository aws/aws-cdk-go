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
type CfnTemplate_HistogramBinOptionsProperty struct {
	// `CfnTemplate.HistogramBinOptionsProperty.BinCount`.
	BinCount interface{} `field:"optional" json:"binCount" yaml:"binCount"`
	// `CfnTemplate.HistogramBinOptionsProperty.BinWidth`.
	BinWidth interface{} `field:"optional" json:"binWidth" yaml:"binWidth"`
	// `CfnTemplate.HistogramBinOptionsProperty.SelectedBinType`.
	SelectedBinType *string `field:"optional" json:"selectedBinType" yaml:"selectedBinType"`
	// `CfnTemplate.HistogramBinOptionsProperty.StartValue`.
	StartValue *float64 `field:"optional" json:"startValue" yaml:"startValue"`
}

