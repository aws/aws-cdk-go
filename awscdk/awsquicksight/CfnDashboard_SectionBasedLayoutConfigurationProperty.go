package awsquicksight


// The configuration for a section-based layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionBasedLayoutConfigurationProperty := &SectionBasedLayoutConfigurationProperty{
//   	BodySections: []interface{}{
//   		&BodySectionConfigurationProperty{
//   			Content: &BodySectionContentProperty{
//   				Layout: &SectionLayoutConfigurationProperty{
//   					FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   						Elements: []interface{}{
//   							&FreeFormLayoutElementProperty{
//   								ElementId: jsii.String("elementId"),
//   								ElementType: jsii.String("elementType"),
//   								Height: jsii.String("height"),
//   								Width: jsii.String("width"),
//   								XAxisLocation: jsii.String("xAxisLocation"),
//   								YAxisLocation: jsii.String("yAxisLocation"),
//
//   								// the properties below are optional
//   								BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   									Color: jsii.String("color"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   								BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   									Color: jsii.String("color"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   								LoadingAnimation: &LoadingAnimationProperty{
//   									Visibility: jsii.String("visibility"),
//   								},
//   								RenderingRules: []interface{}{
//   									&SheetElementRenderingRuleProperty{
//   										ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   											Visibility: jsii.String("visibility"),
//   										},
//   										Expression: jsii.String("expression"),
//   									},
//   								},
//   								SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   									Color: jsii.String("color"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
//
//   			// the properties below are optional
//   			PageBreakConfiguration: &SectionPageBreakConfigurationProperty{
//   				After: &SectionAfterPageBreakProperty{
//   					Status: jsii.String("status"),
//   				},
//   			},
//   			Style: &SectionStyleProperty{
//   				Height: jsii.String("height"),
//   				Padding: &SpacingProperty{
//   					Bottom: jsii.String("bottom"),
//   					Left: jsii.String("left"),
//   					Right: jsii.String("right"),
//   					Top: jsii.String("top"),
//   				},
//   			},
//   		},
//   	},
//   	CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   		PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   			PaperMargin: &SpacingProperty{
//   				Bottom: jsii.String("bottom"),
//   				Left: jsii.String("left"),
//   				Right: jsii.String("right"),
//   				Top: jsii.String("top"),
//   			},
//   			PaperOrientation: jsii.String("paperOrientation"),
//   			PaperSize: jsii.String("paperSize"),
//   		},
//   	},
//   	FooterSections: []interface{}{
//   		&HeaderFooterSectionConfigurationProperty{
//   			Layout: &SectionLayoutConfigurationProperty{
//   				FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   					Elements: []interface{}{
//   						&FreeFormLayoutElementProperty{
//   							ElementId: jsii.String("elementId"),
//   							ElementType: jsii.String("elementType"),
//   							Height: jsii.String("height"),
//   							Width: jsii.String("width"),
//   							XAxisLocation: jsii.String("xAxisLocation"),
//   							YAxisLocation: jsii.String("yAxisLocation"),
//
//   							// the properties below are optional
//   							BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							LoadingAnimation: &LoadingAnimationProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							RenderingRules: []interface{}{
//   								&SheetElementRenderingRuleProperty{
//   									ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   										Visibility: jsii.String("visibility"),
//   									},
//   									Expression: jsii.String("expression"),
//   								},
//   							},
//   							SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
//
//   			// the properties below are optional
//   			Style: &SectionStyleProperty{
//   				Height: jsii.String("height"),
//   				Padding: &SpacingProperty{
//   					Bottom: jsii.String("bottom"),
//   					Left: jsii.String("left"),
//   					Right: jsii.String("right"),
//   					Top: jsii.String("top"),
//   				},
//   			},
//   		},
//   	},
//   	HeaderSections: []interface{}{
//   		&HeaderFooterSectionConfigurationProperty{
//   			Layout: &SectionLayoutConfigurationProperty{
//   				FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   					Elements: []interface{}{
//   						&FreeFormLayoutElementProperty{
//   							ElementId: jsii.String("elementId"),
//   							ElementType: jsii.String("elementType"),
//   							Height: jsii.String("height"),
//   							Width: jsii.String("width"),
//   							XAxisLocation: jsii.String("xAxisLocation"),
//   							YAxisLocation: jsii.String("yAxisLocation"),
//
//   							// the properties below are optional
//   							BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							LoadingAnimation: &LoadingAnimationProperty{
//   								Visibility: jsii.String("visibility"),
//   							},
//   							RenderingRules: []interface{}{
//   								&SheetElementRenderingRuleProperty{
//   									ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   										Visibility: jsii.String("visibility"),
//   									},
//   									Expression: jsii.String("expression"),
//   								},
//   							},
//   							SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
//
//   			// the properties below are optional
//   			Style: &SectionStyleProperty{
//   				Height: jsii.String("height"),
//   				Padding: &SpacingProperty{
//   					Bottom: jsii.String("bottom"),
//   					Left: jsii.String("left"),
//   					Right: jsii.String("right"),
//   					Top: jsii.String("top"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html
//
type CfnDashboard_SectionBasedLayoutConfigurationProperty struct {
	// A list of body section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-bodysections
	//
	BodySections interface{} `field:"required" json:"bodySections" yaml:"bodySections"`
	// The options for the canvas of a section-based layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
	// A list of footer section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-footersections
	//
	FooterSections interface{} `field:"required" json:"footerSections" yaml:"footerSections"`
	// A list of header section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-headersections
	//
	HeaderSections interface{} `field:"required" json:"headerSections" yaml:"headerSections"`
}

