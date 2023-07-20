package awsquicksight


// The configuration of a body section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodySectionConfigurationProperty := &BodySectionConfigurationProperty{
//   	Content: &BodySectionContentProperty{
//   		Layout: &SectionLayoutConfigurationProperty{
//   			FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   				Elements: []interface{}{
//   					&FreeFormLayoutElementProperty{
//   						ElementId: jsii.String("elementId"),
//   						ElementType: jsii.String("elementType"),
//   						Height: jsii.String("height"),
//   						Width: jsii.String("width"),
//   						XAxisLocation: jsii.String("xAxisLocation"),
//   						YAxisLocation: jsii.String("yAxisLocation"),
//
//   						// the properties below are optional
//   						BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   							Color: jsii.String("color"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   							Color: jsii.String("color"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						LoadingAnimation: &LoadingAnimationProperty{
//   							Visibility: jsii.String("visibility"),
//   						},
//   						RenderingRules: []interface{}{
//   							&SheetElementRenderingRuleProperty{
//   								ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   									Visibility: jsii.String("visibility"),
//   								},
//   								Expression: jsii.String("expression"),
//   							},
//   						},
//   						SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   							Color: jsii.String("color"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	SectionId: jsii.String("sectionId"),
//
//   	// the properties below are optional
//   	PageBreakConfiguration: &SectionPageBreakConfigurationProperty{
//   		After: &SectionAfterPageBreakProperty{
//   			Status: jsii.String("status"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html
//
type CfnAnalysis_BodySectionConfigurationProperty struct {
	// The configuration of content in a body section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The unique identifier of a body section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-sectionid
	//
	SectionId *string `field:"required" json:"sectionId" yaml:"sectionId"`
	// The configuration of a page break for a section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-pagebreakconfiguration
	//
	PageBreakConfiguration interface{} `field:"optional" json:"pageBreakConfiguration" yaml:"pageBreakConfiguration"`
	// The style options of a body section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
}

