package previewawsquicksightmixins


// The configuration of a header or footer section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   headerFooterSectionConfigurationProperty := &HeaderFooterSectionConfigurationProperty{
//   	Layout: &SectionLayoutConfigurationProperty{
//   		FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   			Elements: []interface{}{
//   				&FreeFormLayoutElementProperty{
//   					BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   						Color: jsii.String("color"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   						Color: jsii.String("color"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   					ElementId: jsii.String("elementId"),
//   					ElementType: jsii.String("elementType"),
//   					Height: jsii.String("height"),
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
//   					Width: jsii.String("width"),
//   					XAxisLocation: jsii.String("xAxisLocation"),
//   					YAxisLocation: jsii.String("yAxisLocation"),
//   				},
//   			},
//   		},
//   	},
//   	SectionId: jsii.String("sectionId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-headerfootersectionconfiguration.html
//
type CfnAnalysisPropsMixin_HeaderFooterSectionConfigurationProperty struct {
	// The layout configuration of the header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-headerfootersectionconfiguration.html#cfn-quicksight-analysis-headerfootersectionconfiguration-layout
	//
	Layout interface{} `field:"optional" json:"layout" yaml:"layout"`
	// The unique identifier of the header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-headerfootersectionconfiguration.html#cfn-quicksight-analysis-headerfootersectionconfiguration-sectionid
	//
	SectionId *string `field:"optional" json:"sectionId" yaml:"sectionId"`
	// The style options of a header or footer section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-headerfootersectionconfiguration.html#cfn-quicksight-analysis-headerfootersectionconfiguration-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
}

