package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimePickerControlDisplayOptionsProperty := &DateTimePickerControlDisplayOptionsProperty{
//   	DateTimeFormat: jsii.String("dateTimeFormat"),
//   	TitleOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnTemplate_DateTimePickerControlDisplayOptionsProperty struct {
	// `CfnTemplate.DateTimePickerControlDisplayOptionsProperty.DateTimeFormat`.
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// `CfnTemplate.DateTimePickerControlDisplayOptionsProperty.TitleOptions`.
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

