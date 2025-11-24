package mixinsawsquicksight


// The configuration of a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   freeFormLayoutConfigurationProperty := &FreeFormLayoutConfigurationProperty{
//   	CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		},
//   	},
//   	Elements: []interface{}{
//   		&FreeFormLayoutElementProperty{
//   			BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   				Color: jsii.String("color"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   				Color: jsii.String("color"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			ElementId: jsii.String("elementId"),
//   			ElementType: jsii.String("elementType"),
//   			Height: jsii.String("height"),
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
//   			Width: jsii.String("width"),
//   			XAxisLocation: jsii.String("xAxisLocation"),
//   			YAxisLocation: jsii.String("yAxisLocation"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutconfiguration.html
//
type CfnTemplatePropsMixin_FreeFormLayoutConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutconfiguration.html#cfn-quicksight-template-freeformlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
	// The elements that are included in a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutconfiguration.html#cfn-quicksight-template-freeformlayoutconfiguration-elements
	//
	Elements interface{} `field:"optional" json:"elements" yaml:"elements"`
}

