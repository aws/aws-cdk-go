package awsquicksight


// The configuration of a header or footer section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerFooterSectionConfigurationProperty := &HeaderFooterSectionConfigurationProperty{
//   	Layout: &SectionLayoutConfigurationProperty{
//   		FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   			Elements: []interface{}{
//   				&FreeFormLayoutElementProperty{
//   					ElementId: jsii.String("elementId"),
//   					ElementType: jsii.String("elementType"),
//   					Height: jsii.String("height"),
//   					Width: jsii.String("width"),
//   					XAxisLocation: jsii.String("xAxisLocation"),
//   					YAxisLocation: jsii.String("yAxisLocation"),
//
//   					// the properties below are optional
//   					BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   						Color: jsii.String("color"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   						Color: jsii.String("color"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					LoadingAnimation: &LoadingAnimationProperty{
//   						Visibility: jsii.String("visibility"),
//   					},
//   					RenderingRules: []interface{}{
//   						&SheetElementRenderingRuleProperty{
//   							ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   					},
//   					SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   						Color: jsii.String("color"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					Visibility: jsii.String("visibility"),
//   				},
//   			},
//   		},
//   	},
//   	SectionId: jsii.String("sectionId"),
//
//   	// the properties below are optional
//   	Style: &SectionStyleProperty{
//   		Height: jsii.String("height"),
//   		Padding: &SpacingProperty{
//   			Bottom: jsii.String("bottom"),
//   			Left: jsii.String("left"),
//   			Right: jsii.String("right"),
//   			Top: jsii.String("top"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-headerfootersectionconfiguration.html
//
type CfnDashboard_HeaderFooterSectionConfigurationProperty struct {
	// The layout configuration of the header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-headerfootersectionconfiguration.html#cfn-quicksight-dashboard-headerfootersectionconfiguration-layout
	//
	Layout interface{} `field:"required" json:"layout" yaml:"layout"`
	// The unique identifier of the header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-headerfootersectionconfiguration.html#cfn-quicksight-dashboard-headerfootersectionconfiguration-sectionid
	//
	SectionId *string `field:"required" json:"sectionId" yaml:"sectionId"`
	// The style options of a header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-headerfootersectionconfiguration.html#cfn-quicksight-dashboard-headerfootersectionconfiguration-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
}

