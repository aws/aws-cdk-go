package previewawsquicksightmixins


// The configuration for a section-based layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sectionBasedLayoutConfigurationProperty := &SectionBasedLayoutConfigurationProperty{
//   	BodySections: []interface{}{
//   		&BodySectionConfigurationProperty{
//   			Content: &BodySectionContentProperty{
//   				Layout: &SectionLayoutConfigurationProperty{
//   					FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   						Elements: []interface{}{
//   							&FreeFormLayoutElementProperty{
//   								BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   									Color: jsii.String("color"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   								BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   									Color: jsii.String("color"),
//   									Visibility: jsii.String("visibility"),
//   								},
//   								ElementId: jsii.String("elementId"),
//   								ElementType: jsii.String("elementType"),
//   								Height: jsii.String("height"),
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
//   								Width: jsii.String("width"),
//   								XAxisLocation: jsii.String("xAxisLocation"),
//   								YAxisLocation: jsii.String("yAxisLocation"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			PageBreakConfiguration: &SectionPageBreakConfigurationProperty{
//   				After: &SectionAfterPageBreakProperty{
//   					Status: jsii.String("status"),
//   				},
//   			},
//   			RepeatConfiguration: &BodySectionRepeatConfigurationProperty{
//   				DimensionConfigurations: []interface{}{
//   					&BodySectionRepeatDimensionConfigurationProperty{
//   						DynamicCategoryDimensionConfiguration: &BodySectionDynamicCategoryDimensionConfigurationProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							Limit: jsii.Number(123),
//   							SortByMetrics: []interface{}{
//   								&ColumnSortProperty{
//   									AggregationFunction: &AggregationFunctionProperty{
//   										AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   											SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   											ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   										},
//   										CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   										DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   										NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   											PercentileAggregation: &PercentileAggregationProperty{
//   												PercentileValue: jsii.Number(123),
//   											},
//   											SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   										},
//   									},
//   									Direction: jsii.String("direction"),
//   									SortBy: &ColumnIdentifierProperty{
//   										ColumnName: jsii.String("columnName"),
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									},
//   								},
//   							},
//   						},
//   						DynamicNumericDimensionConfiguration: &BodySectionDynamicNumericDimensionConfigurationProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							Limit: jsii.Number(123),
//   							SortByMetrics: []interface{}{
//   								&ColumnSortProperty{
//   									AggregationFunction: &AggregationFunctionProperty{
//   										AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   											SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   											ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   										},
//   										CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   										DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   										NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   											PercentileAggregation: &PercentileAggregationProperty{
//   												PercentileValue: jsii.Number(123),
//   											},
//   											SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   										},
//   									},
//   									Direction: jsii.String("direction"),
//   									SortBy: &ColumnIdentifierProperty{
//   										ColumnName: jsii.String("columnName"),
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				NonRepeatingVisuals: []*string{
//   					jsii.String("nonRepeatingVisuals"),
//   				},
//   				PageBreakConfiguration: &BodySectionRepeatPageBreakConfigurationProperty{
//   					After: &SectionAfterPageBreakProperty{
//   						Status: jsii.String("status"),
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
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
//   							BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							ElementId: jsii.String("elementId"),
//   							ElementType: jsii.String("elementType"),
//   							Height: jsii.String("height"),
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
//   							Width: jsii.String("width"),
//   							XAxisLocation: jsii.String("xAxisLocation"),
//   							YAxisLocation: jsii.String("yAxisLocation"),
//   						},
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
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
//   							BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   								Color: jsii.String("color"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   							ElementId: jsii.String("elementId"),
//   							ElementType: jsii.String("elementType"),
//   							Height: jsii.String("height"),
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
//   							Width: jsii.String("width"),
//   							XAxisLocation: jsii.String("xAxisLocation"),
//   							YAxisLocation: jsii.String("yAxisLocation"),
//   						},
//   					},
//   				},
//   			},
//   			SectionId: jsii.String("sectionId"),
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
type CfnDashboardPropsMixin_SectionBasedLayoutConfigurationProperty struct {
	// A list of body section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-bodysections
	//
	BodySections interface{} `field:"optional" json:"bodySections" yaml:"bodySections"`
	// The options for the canvas of a section-based layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
	// A list of footer section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-footersections
	//
	FooterSections interface{} `field:"optional" json:"footerSections" yaml:"footerSections"`
	// A list of header section configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutconfiguration.html#cfn-quicksight-dashboard-sectionbasedlayoutconfiguration-headersections
	//
	HeaderSections interface{} `field:"optional" json:"headerSections" yaml:"headerSections"`
}

