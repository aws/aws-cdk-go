package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionLayoutConfigurationProperty := &SectionLayoutConfigurationProperty{
//   	FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   		Elements: []interface{}{
//   			&FreeFormLayoutElementProperty{
//   				ElementId: jsii.String("elementId"),
//   				ElementType: jsii.String("elementType"),
//   				Height: jsii.String("height"),
//   				Width: jsii.String("width"),
//   				XAxisLocation: jsii.String("xAxisLocation"),
//   				YAxisLocation: jsii.String("yAxisLocation"),
//
//   				// the properties below are optional
//   				BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				LoadingAnimation: &LoadingAnimationProperty{
//   					Visibility: jsii.String("visibility"),
//   				},
//   				RenderingRules: []interface{}{
//   					&SheetElementRenderingRuleProperty{
//   						ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   				SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_SectionLayoutConfigurationProperty struct {
	// `CfnTemplate.SectionLayoutConfigurationProperty.FreeFormLayout`.
	FreeFormLayout interface{} `field:"required" json:"freeFormLayout" yaml:"freeFormLayout"`
}

