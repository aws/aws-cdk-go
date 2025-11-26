package previewawsquicksightmixins


// An element within a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   freeFormLayoutElementProperty := &FreeFormLayoutElementProperty{
//   	BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	ElementId: jsii.String("elementId"),
//   	ElementType: jsii.String("elementType"),
//   	Height: jsii.String("height"),
//   	LoadingAnimation: &LoadingAnimationProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	RenderingRules: []interface{}{
//   		&SheetElementRenderingRuleProperty{
//   			ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   	Width: jsii.String("width"),
//   	XAxisLocation: jsii.String("xAxisLocation"),
//   	YAxisLocation: jsii.String("yAxisLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html
//
type CfnTemplatePropsMixin_FreeFormLayoutElementProperty struct {
	// The background style configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-backgroundstyle
	//
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// The border style configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-borderstyle
	//
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// A unique identifier for an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-elementid
	//
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// The type of element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-elementtype
	//
	ElementType *string `field:"optional" json:"elementType" yaml:"elementType"`
	// The height of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-height
	//
	Height *string `field:"optional" json:"height" yaml:"height"`
	// The loading animation configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-loadinganimation
	//
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// The rendering rules that determine when an element should be displayed within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-renderingrules
	//
	RenderingRules interface{} `field:"optional" json:"renderingRules" yaml:"renderingRules"`
	// The border style configuration of a free-form layout element.
	//
	// This border style is used when the element is selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-selectedborderstyle
	//
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
	// The visibility of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The width of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
	// The x-axis coordinate of the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-xaxislocation
	//
	XAxisLocation *string `field:"optional" json:"xAxisLocation" yaml:"xAxisLocation"`
	// The y-axis coordinate of the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutelement.html#cfn-quicksight-template-freeformlayoutelement-yaxislocation
	//
	YAxisLocation *string `field:"optional" json:"yAxisLocation" yaml:"yAxisLocation"`
}

