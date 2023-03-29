package awsquicksight


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
	// `CfnAnalysis.FreeFormLayoutElementProperty.ElementId`.
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.ElementType`.
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.Height`.
	Height *string `field:"required" json:"height" yaml:"height"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.Width`.
	Width *string `field:"required" json:"width" yaml:"width"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.XAxisLocation`.
	XAxisLocation *string `field:"required" json:"xAxisLocation" yaml:"xAxisLocation"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.YAxisLocation`.
	YAxisLocation *string `field:"required" json:"yAxisLocation" yaml:"yAxisLocation"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.BackgroundStyle`.
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.BorderStyle`.
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.LoadingAnimation`.
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.RenderingRules`.
	RenderingRules interface{} `field:"optional" json:"renderingRules" yaml:"renderingRules"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.SelectedBorderStyle`.
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
	// `CfnAnalysis.FreeFormLayoutElementProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

