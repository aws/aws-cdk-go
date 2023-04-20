package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   funnelChartDataLabelOptionsProperty := &FunnelChartDataLabelOptionsProperty{
//   	CategoryLabelVisibility: jsii.String("categoryLabelVisibility"),
//   	LabelColor: jsii.String("labelColor"),
//   	LabelFontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	MeasureDataLabelStyle: jsii.String("measureDataLabelStyle"),
//   	MeasureLabelVisibility: jsii.String("measureLabelVisibility"),
//   	Position: jsii.String("position"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_FunnelChartDataLabelOptionsProperty struct {
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.CategoryLabelVisibility`.
	CategoryLabelVisibility *string `field:"optional" json:"categoryLabelVisibility" yaml:"categoryLabelVisibility"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.LabelColor`.
	LabelColor *string `field:"optional" json:"labelColor" yaml:"labelColor"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.LabelFontConfiguration`.
	LabelFontConfiguration interface{} `field:"optional" json:"labelFontConfiguration" yaml:"labelFontConfiguration"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.MeasureDataLabelStyle`.
	MeasureDataLabelStyle *string `field:"optional" json:"measureDataLabelStyle" yaml:"measureDataLabelStyle"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.MeasureLabelVisibility`.
	MeasureLabelVisibility *string `field:"optional" json:"measureLabelVisibility" yaml:"measureLabelVisibility"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.Position`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// `CfnAnalysis.FunnelChartDataLabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

