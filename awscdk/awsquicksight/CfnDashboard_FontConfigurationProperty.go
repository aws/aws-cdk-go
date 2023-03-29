package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontConfigurationProperty := &FontConfigurationProperty{
//   	FontColor: jsii.String("fontColor"),
//   	FontDecoration: jsii.String("fontDecoration"),
//   	FontSize: &FontSizeProperty{
//   		Relative: jsii.String("relative"),
//   	},
//   	FontStyle: jsii.String("fontStyle"),
//   	FontWeight: &FontWeightProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
type CfnDashboard_FontConfigurationProperty struct {
	// `CfnDashboard.FontConfigurationProperty.FontColor`.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// `CfnDashboard.FontConfigurationProperty.FontDecoration`.
	FontDecoration *string `field:"optional" json:"fontDecoration" yaml:"fontDecoration"`
	// `CfnDashboard.FontConfigurationProperty.FontSize`.
	FontSize interface{} `field:"optional" json:"fontSize" yaml:"fontSize"`
	// `CfnDashboard.FontConfigurationProperty.FontStyle`.
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// `CfnDashboard.FontConfigurationProperty.FontWeight`.
	FontWeight interface{} `field:"optional" json:"fontWeight" yaml:"fontWeight"`
}

