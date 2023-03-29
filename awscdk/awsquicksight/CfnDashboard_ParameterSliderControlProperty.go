package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterSliderControlProperty := &ParameterSliderControlProperty{
//   	MaximumValue: jsii.Number(123),
//   	MinimumValue: jsii.Number(123),
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   	StepSize: jsii.Number(123),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &SliderControlDisplayOptionsProperty{
//   		TitleOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnDashboard_ParameterSliderControlProperty struct {
	// `CfnDashboard.ParameterSliderControlProperty.MaximumValue`.
	MaximumValue *float64 `field:"required" json:"maximumValue" yaml:"maximumValue"`
	// `CfnDashboard.ParameterSliderControlProperty.MinimumValue`.
	MinimumValue *float64 `field:"required" json:"minimumValue" yaml:"minimumValue"`
	// `CfnDashboard.ParameterSliderControlProperty.ParameterControlId`.
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// `CfnDashboard.ParameterSliderControlProperty.SourceParameterName`.
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// `CfnDashboard.ParameterSliderControlProperty.StepSize`.
	StepSize *float64 `field:"required" json:"stepSize" yaml:"stepSize"`
	// `CfnDashboard.ParameterSliderControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnDashboard.ParameterSliderControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

