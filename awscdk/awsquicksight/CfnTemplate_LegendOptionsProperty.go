package awsquicksight


// The options for the legend setup of a visual.
//
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
type CfnTemplate_LegendOptionsProperty struct {
	// The height of the legend.
	//
	// If this value is omitted, a default height is used when rendering.
	Height *string `field:"optional" json:"height" yaml:"height"`
	// The positions for the legend. Choose one of the following options:.
	//
	// - `AUTO`
	// - `RIGHT`
	// - `BOTTOM`
	// - `LEFT`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// The custom title for the legend.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// Determines whether or not the legend is visible.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The width of the legend.
	//
	// If this value is omitted, a default width is used when rendering.
	Width *string `field:"optional" json:"width" yaml:"width"`
}

