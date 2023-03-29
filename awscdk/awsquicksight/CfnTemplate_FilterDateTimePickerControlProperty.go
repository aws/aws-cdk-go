package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDateTimePickerControlProperty := &FilterDateTimePickerControlProperty{
//   	FilterControlId: jsii.String("filterControlId"),
//   	SourceFilterId: jsii.String("sourceFilterId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
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
type CfnTemplate_FilterDateTimePickerControlProperty struct {
	// `CfnTemplate.FilterDateTimePickerControlProperty.FilterControlId`.
	FilterControlId *string `field:"required" json:"filterControlId" yaml:"filterControlId"`
	// `CfnTemplate.FilterDateTimePickerControlProperty.SourceFilterId`.
	SourceFilterId *string `field:"required" json:"sourceFilterId" yaml:"sourceFilterId"`
	// `CfnTemplate.FilterDateTimePickerControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnTemplate.FilterDateTimePickerControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
	// `CfnTemplate.FilterDateTimePickerControlProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

