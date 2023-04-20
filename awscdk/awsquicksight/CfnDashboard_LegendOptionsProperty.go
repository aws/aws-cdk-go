package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   legendOptionsProperty := &LegendOptionsProperty{
//   	Height: jsii.String("height"),
//   	Position: jsii.String("position"),
//   	Title: &LabelOptionsProperty{
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
//   	Visibility: jsii.String("visibility"),
//   	Width: jsii.String("width"),
//   }
//
type CfnDashboard_LegendOptionsProperty struct {
	// `CfnDashboard.LegendOptionsProperty.Height`.
	Height *string `field:"optional" json:"height" yaml:"height"`
	// `CfnDashboard.LegendOptionsProperty.Position`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// `CfnDashboard.LegendOptionsProperty.Title`.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// `CfnDashboard.LegendOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// `CfnDashboard.LegendOptionsProperty.Width`.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

