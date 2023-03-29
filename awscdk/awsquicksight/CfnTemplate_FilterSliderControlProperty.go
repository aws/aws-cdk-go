package awsquicksight


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
	// `CfnTemplate.FilterSliderControlProperty.FilterControlId`.
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// `CfnTemplate.FilterSliderControlProperty.MaximumValue`.
	MaximumValue *float64 `field:"required" json:"maximumValue" yaml:"maximumValue"`
	// `CfnTemplate.FilterSliderControlProperty.MinimumValue`.
	MinimumValue *float64 `field:"required" json:"minimumValue" yaml:"minimumValue"`
	// `CfnTemplate.FilterSliderControlProperty.SourceFilterId`.
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// `CfnTemplate.FilterSliderControlProperty.StepSize`.
	StepSize *float64 `field:"required" json:"stepSize" yaml:"stepSize"`
	// `CfnTemplate.FilterSliderControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnTemplate.FilterSliderControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// `CfnTemplate.FilterSliderControlProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

