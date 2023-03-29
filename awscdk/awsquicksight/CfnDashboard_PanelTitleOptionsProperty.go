package awsquicksight


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
type CfnDashboard_PanelTitleOptionsProperty struct {
	// `CfnDashboard.PanelTitleOptionsProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// `CfnDashboard.PanelTitleOptionsProperty.HorizontalTextAlignment`.
	HorizontalTextAlignment *string `field:"optional" json:"horizontalTextAlignment" yaml:"horizontalTextAlignment"`
	// `CfnDashboard.PanelTitleOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

