package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterDateTimePickerControlProperty := &ParameterDateTimePickerControlProperty{
//   	ParameterControlId: jsii.String("parameterControlId"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
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
//   }
//
type CfnTemplate_ParameterDateTimePickerControlProperty struct {
	// `CfnTemplate.ParameterDateTimePickerControlProperty.ParameterControlId`.
	ParameterControlId *string `field:"required" json:"parameterControlId" yaml:"parameterControlId"`
	// `CfnTemplate.ParameterDateTimePickerControlProperty.SourceParameterName`.
	SourceParameterName *string `field:"required" json:"sourceParameterName" yaml:"sourceParameterName"`
	// `CfnTemplate.ParameterDateTimePickerControlProperty.Title`.
	Title *string `field:"required" json:"title" yaml:"title"`
	// `CfnTemplate.ParameterDateTimePickerControlProperty.DisplayOptions`.
	DisplayOptions interface{} `field:"optional" json:"displayOptions" yaml:"displayOptions"`
}

