package awsquicksight


// A control to display a horizontal toggle bar.
//
// This is used to change a value by sliding the toggle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterSliderControlProperty := &FilterSliderControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	MaximumValue: jsii.Number(123),
//   	MinimumValue: jsii.Number(123),
//   	SourceFilterId: jsii.String("sourceFilterId"),
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
//   	Type: jsii.String("type"),
//   }
//
type CfnTemplate_FilterSliderControlProperty struct {
	// The ID of the `FilterSliderControl` .
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// The smaller value that is displayed at the left of the slider.
	MaximumValue *float64 `field:"required" json:"maximumValue" yaml:"maximumValue"`
	// The larger value that is displayed at the right of the slider.
	MinimumValue *float64 `field:"required" json:"minimumValue" yaml:"minimumValue"`
	// The source filter ID of the `FilterSliderControl` .
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// The number of increments that the slider bar is divided into.
	StepSize *float64 `field:"required" json:"stepSize" yaml:"stepSize"`
	// The title of the `FilterSliderControl` .
	Title *string `field:"required" json:"title" yaml:"title"`
	// The display options of a control.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// The type of `FilterSliderControl` . Choose one of the following options:.
	//
	// - `SINGLE_POINT` : Filter against(equals) a single data point.
	// - `RANGE` : Filter data that is in a specified range.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

