package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   panelConfigurationProperty := &PanelConfigurationProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BackgroundVisibility: jsii.String("backgroundVisibility"),
//   	BorderColor: jsii.String("borderColor"),
//   	BorderStyle: jsii.String("borderStyle"),
//   	BorderThickness: jsii.String("borderThickness"),
//   	BorderVisibility: jsii.String("borderVisibility"),
//   	GutterSpacing: jsii.String("gutterSpacing"),
//   	GutterVisibility: jsii.String("gutterVisibility"),
//   	Title: &PanelTitleOptionsProperty{
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
//   		HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnAnalysis_PanelConfigurationProperty struct {
	// `CfnAnalysis.PanelConfigurationProperty.BackgroundColor`.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// `CfnAnalysis.PanelConfigurationProperty.BackgroundVisibility`.
	BackgroundVisibility *string `field:"optional" json:"backgroundVisibility" yaml:"backgroundVisibility"`
	// `CfnAnalysis.PanelConfigurationProperty.BorderColor`.
	BorderColor *string `field:"optional" json:"borderColor" yaml:"borderColor"`
	// `CfnAnalysis.PanelConfigurationProperty.BorderStyle`.
	BorderStyle *string `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// `CfnAnalysis.PanelConfigurationProperty.BorderThickness`.
	BorderThickness *string `field:"optional" json:"borderThickness" yaml:"borderThickness"`
	// `CfnAnalysis.PanelConfigurationProperty.BorderVisibility`.
	BorderVisibility *string `field:"optional" json:"borderVisibility" yaml:"borderVisibility"`
	// `CfnAnalysis.PanelConfigurationProperty.GutterSpacing`.
	GutterSpacing *string `field:"optional" json:"gutterSpacing" yaml:"gutterSpacing"`
	// `CfnAnalysis.PanelConfigurationProperty.GutterVisibility`.
	GutterVisibility *string `field:"optional" json:"gutterVisibility" yaml:"gutterVisibility"`
	// `CfnAnalysis.PanelConfigurationProperty.Title`.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}

