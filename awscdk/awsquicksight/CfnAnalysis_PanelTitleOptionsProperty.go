package awsquicksight


// The options that determine the title styles for each small multiples panel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   panelTitleOptionsProperty := &PanelTitleOptionsProperty{
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
//   	HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_PanelTitleOptionsProperty struct {
	// `CfnAnalysis.PanelTitleOptionsProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// Sets the horizontal text alignment of the title within each panel.
	HorizontalTextAlignment *string `field:"optional" json:"horizontalTextAlignment" yaml:"horizontalTextAlignment"`
	// Determines whether or not panel titles are displayed.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

