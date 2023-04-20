package awsquicksight


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
type CfnAnalysis_LabelOptionsProperty struct {
	// `CfnAnalysis.LabelOptionsProperty.CustomLabel`.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// `CfnAnalysis.LabelOptionsProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// `CfnAnalysis.LabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

