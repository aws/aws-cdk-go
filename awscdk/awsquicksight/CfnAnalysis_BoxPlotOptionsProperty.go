package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   boxPlotOptionsProperty := &BoxPlotOptionsProperty{
//   	AllDataPointsVisibility: jsii.String("allDataPointsVisibility"),
//   	OutlierVisibility: jsii.String("outlierVisibility"),
//   	StyleOptions: &BoxPlotStyleOptionsProperty{
//   		FillStyle: jsii.String("fillStyle"),
//   	},
//   }
//
type CfnAnalysis_BoxPlotOptionsProperty struct {
	// `CfnAnalysis.BoxPlotOptionsProperty.AllDataPointsVisibility`.
	AllDataPointsVisibility *string `field:"optional" json:"allDataPointsVisibility" yaml:"allDataPointsVisibility"`
	// `CfnAnalysis.BoxPlotOptionsProperty.OutlierVisibility`.
	OutlierVisibility *string `field:"optional" json:"outlierVisibility" yaml:"outlierVisibility"`
	// `CfnAnalysis.BoxPlotOptionsProperty.StyleOptions`.
	StyleOptions interface{} `field:"optional" json:"styleOptions" yaml:"styleOptions"`
}

