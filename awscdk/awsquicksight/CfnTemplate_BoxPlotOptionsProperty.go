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
type CfnTemplate_BoxPlotOptionsProperty struct {
	// `CfnTemplate.BoxPlotOptionsProperty.AllDataPointsVisibility`.
	AllDataPointsVisibility *string `field:"optional" json:"allDataPointsVisibility" yaml:"allDataPointsVisibility"`
	// `CfnTemplate.BoxPlotOptionsProperty.OutlierVisibility`.
	OutlierVisibility *string `field:"optional" json:"outlierVisibility" yaml:"outlierVisibility"`
	// `CfnTemplate.BoxPlotOptionsProperty.StyleOptions`.
	StyleOptions interface{} `field:"optional" json:"styleOptions" yaml:"styleOptions"`
}

