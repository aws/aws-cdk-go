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
//   	BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
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
//   }
//
type CfnAnalysis_FreeFormLayoutElementProperty struct {
	// A unique identifier for an element within a free-form layout.
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// The type of element.
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// The height of an element within a free-form layout.
	Height *string `field:"required" json:"height" yaml:"height"`
	// The width of an element within a free-form layout.
	Width *string `field:"required" json:"width" yaml:"width"`
	// The x-axis coordinate of the element.
	XAxisLocation *string `field:"required" json:"xAxisLocation" yaml:"xAxisLocation"`
	// The y-axis coordinate of the element.
	YAxisLocation *string `field:"required" json:"yAxisLocation" yaml:"yAxisLocation"`
	// The background style configuration of a free-form layout element.
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// The border style configuration of a free-form layout element.
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// The loading animation configuration of a free-form layout element.
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// The rendering rules that determine when an element should be displayed within a free-form layout.
	RenderingRules interface{} `field:"optional" json:"renderingRules" yaml:"renderingRules"`
	// The border style configuration of a free-form layout element.
	//
	// This border style is used when the element is selected.
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
	// The visibility of an element within a free-form layout.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

