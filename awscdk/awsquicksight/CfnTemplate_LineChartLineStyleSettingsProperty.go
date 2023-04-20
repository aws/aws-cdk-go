package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartLineStyleSettingsProperty := &LineChartLineStyleSettingsProperty{
//   	LineInterpolation: jsii.String("lineInterpolation"),
//   	LineStyle: jsii.String("lineStyle"),
//   	LineVisibility: jsii.String("lineVisibility"),
//   	LineWidth: jsii.String("lineWidth"),
//   }
//
type CfnTemplate_LineChartLineStyleSettingsProperty struct {
	// `CfnTemplate.LineChartLineStyleSettingsProperty.LineInterpolation`.
	LineInterpolation *string `field:"optional" json:"lineInterpolation" yaml:"lineInterpolation"`
	// `CfnTemplate.LineChartLineStyleSettingsProperty.LineStyle`.
	LineStyle *string `field:"optional" json:"lineStyle" yaml:"lineStyle"`
	// `CfnTemplate.LineChartLineStyleSettingsProperty.LineVisibility`.
	LineVisibility *string `field:"optional" json:"lineVisibility" yaml:"lineVisibility"`
	// `CfnTemplate.LineChartLineStyleSettingsProperty.LineWidth`.
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}

