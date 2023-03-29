package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textFieldControlDisplayOptionsProperty := &TextFieldControlDisplayOptionsProperty{
//   	PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
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
type CfnTemplate_TextFieldControlDisplayOptionsProperty struct {
	// `CfnTemplate.TextFieldControlDisplayOptionsProperty.PlaceholderOptions`.
	PlaceholderOptions interface{} `field:"optional" json:"placeholderOptions" yaml:"placeholderOptions"`
	// `CfnTemplate.TextFieldControlDisplayOptionsProperty.TitleOptions`.
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

