package awsquicksight


// Configures the display properties of the given text.
//
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
type CfnAnalysis_FontConfigurationProperty struct {
	// Determines the color of the text.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Determines the appearance of decorative lines on the text.
	FontDecoration *string `field:"optional" json:"fontDecoration" yaml:"fontDecoration"`
	// The option that determines the text display size.
	FontSize interface{} `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Determines the text display face that is inherited by the given font family.
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// The option that determines the text display weight, or boldness.
	FontWeight interface{} `field:"optional" json:"fontWeight" yaml:"fontWeight"`
}

