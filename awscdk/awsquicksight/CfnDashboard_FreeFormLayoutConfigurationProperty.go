package awsquicksight


// The configuration of a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutConfigurationProperty := &FreeFormLayoutConfigurationProperty{
//   	Elements: []interface{}{
//   		&FreeFormLayoutElementProperty{
//   			ElementId: jsii.String("elementId"),
//   			ElementType: jsii.String("elementType"),
//   			Height: jsii.String("height"),
//   			Width: jsii.String("width"),
//   			XAxisLocation: jsii.String("xAxisLocation"),
//   			YAxisLocation: jsii.String("yAxisLocation"),
//
//   			// the properties below are optional
//   			BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   				Color: jsii.String("color"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   				Color: jsii.String("color"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			LoadingAnimation: &LoadingAnimationProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			RenderingRules: []interface{}{
//   				&SheetElementRenderingRuleProperty{
//   					ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					Expression: jsii.String("expression"),
//   				},
//   			},
//   			SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   				Color: jsii.String("color"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html
//
type CfnDashboard_FreeFormLayoutConfigurationProperty struct {
	// The elements that are included in a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html#cfn-quicksight-dashboard-freeformlayoutconfiguration-elements
	//
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutconfiguration.html#cfn-quicksight-dashboard-freeformlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

