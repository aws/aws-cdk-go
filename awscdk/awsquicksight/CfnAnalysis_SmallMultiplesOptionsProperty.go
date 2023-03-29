package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smallMultiplesOptionsProperty := &SmallMultiplesOptionsProperty{
//   	MaxVisibleColumns: jsii.Number(123),
//   	MaxVisibleRows: jsii.Number(123),
//   	PanelConfiguration: &PanelConfigurationProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		BackgroundVisibility: jsii.String("backgroundVisibility"),
//   		BorderColor: jsii.String("borderColor"),
//   		BorderStyle: jsii.String("borderStyle"),
//   		BorderThickness: jsii.String("borderThickness"),
//   		BorderVisibility: jsii.String("borderVisibility"),
//   		GutterSpacing: jsii.String("gutterSpacing"),
//   		GutterVisibility: jsii.String("gutterVisibility"),
//   		Title: &PanelTitleOptionsProperty{
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnAnalysis_SmallMultiplesOptionsProperty struct {
	// `CfnAnalysis.SmallMultiplesOptionsProperty.MaxVisibleColumns`.
	MaxVisibleColumns *float64 `field:"optional" json:"maxVisibleColumns" yaml:"maxVisibleColumns"`
	// `CfnAnalysis.SmallMultiplesOptionsProperty.MaxVisibleRows`.
	MaxVisibleRows *float64 `field:"optional" json:"maxVisibleRows" yaml:"maxVisibleRows"`
	// `CfnAnalysis.SmallMultiplesOptionsProperty.PanelConfiguration`.
	PanelConfiguration interface{} `field:"optional" json:"panelConfiguration" yaml:"panelConfiguration"`
}

