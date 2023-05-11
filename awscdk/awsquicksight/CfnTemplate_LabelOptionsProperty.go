package awsquicksight


// The share label options for the labels.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelOptionsProperty := &LabelOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	FontConfiguration: &FontConfigurationProperty{
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
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnTemplate_LabelOptionsProperty struct {
	// The text for the label.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The font configuration of the label.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// Determines whether or not the label is visible.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

