package awsquicksight


// An element within a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutElementProperty := &FreeFormLayoutElementProperty{
//   	ElementId: jsii.String("elementId"),
//   	ElementType: jsii.String("elementType"),
//   	Height: jsii.String("height"),
//   	Width: jsii.String("width"),
//   	XAxisLocation: jsii.String("xAxisLocation"),
//   	YAxisLocation: jsii.String("yAxisLocation"),
//
//   	// the properties below are optional
//   	BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	BorderRadius: jsii.String("borderRadius"),
//   	BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   		Width: jsii.String("width"),
//   	},
//   	LoadingAnimation: &LoadingAnimationProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Padding: jsii.String("padding"),
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
//   		Width: jsii.String("width"),
//   	},
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html
//
type CfnDashboard_FreeFormLayoutElementProperty struct {
	// A unique identifier for an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-elementid
	//
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// The type of element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-elementtype
	//
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// The height of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-height
	//
	Height *string `field:"required" json:"height" yaml:"height"`
	// The width of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-width
	//
	Width *string `field:"required" json:"width" yaml:"width"`
	// The x-axis coordinate of the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-xaxislocation
	//
	XAxisLocation *string `field:"required" json:"xAxisLocation" yaml:"xAxisLocation"`
	// The y-axis coordinate of the element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-yaxislocation
	//
	YAxisLocation *string `field:"required" json:"yAxisLocation" yaml:"yAxisLocation"`
	// The background style configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-backgroundstyle
	//
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// The border radius of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-borderradius
	//
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// The border style configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-borderstyle
	//
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// The loading animation configuration of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-loadinganimation
	//
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// The padding of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-padding
	//
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
	// The rendering rules that determine when an element should be displayed within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-renderingrules
	//
	RenderingRules interface{} `field:"optional" json:"renderingRules" yaml:"renderingRules"`
	// The border style configuration of a free-form layout element.
	//
	// This border style is used when the element is selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-selectedborderstyle
	//
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
	// The visibility of an element within a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelement.html#cfn-quicksight-dashboard-freeformlayoutelement-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

