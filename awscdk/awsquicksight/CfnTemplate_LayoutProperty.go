package awsquicksight


// A `Layout` defines the placement of elements within a sheet.
//
// For more information, see [Types of layout](https://docs.aws.amazon.com/quicksight/latest/user/types-of-layout.html) in the *Amazon Quick Suite User Guide* .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   layoutProperty := &LayoutProperty{
//   	Configuration: &LayoutConfigurationProperty{
//   		FreeFormLayout: &FreeFormLayoutConfigurationProperty{
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
//
//   			// the properties below are optional
//   			CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   				ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				},
//   			},
//   		},
//   		GridLayout: &GridLayoutConfigurationProperty{
//   			Elements: []interface{}{
//   				&GridLayoutElementProperty{
//   					ColumnSpan: jsii.Number(123),
//   					ElementId: jsii.String("elementId"),
//   					ElementType: jsii.String("elementType"),
//   					RowSpan: jsii.Number(123),
//
//   					// the properties below are optional
//   					ColumnIndex: jsii.Number(123),
//   					RowIndex: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   				ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   					ResizeOption: jsii.String("resizeOption"),
//
//   					// the properties below are optional
//   					OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				},
//   			},
//   		},
//   		SectionBasedLayout: &SectionBasedLayoutConfigurationProperty{
//   			BodySections: []interface{}{
//   				&BodySectionConfigurationProperty{
//   					Content: &BodySectionContentProperty{
//   						Layout: &SectionLayoutConfigurationProperty{
//   							FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   								Elements: []interface{}{
//   									&FreeFormLayoutElementProperty{
//   										ElementId: jsii.String("elementId"),
//   										ElementType: jsii.String("elementType"),
//   										Height: jsii.String("height"),
//   										Width: jsii.String("width"),
//   										XAxisLocation: jsii.String("xAxisLocation"),
//   										YAxisLocation: jsii.String("yAxisLocation"),
//
//   										// the properties below are optional
//   										BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   											Color: jsii.String("color"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   										BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   											Color: jsii.String("color"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   										LoadingAnimation: &LoadingAnimationProperty{
//   											Visibility: jsii.String("visibility"),
//   										},
//   										RenderingRules: []interface{}{
//   											&SheetElementRenderingRuleProperty{
//   												ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   													Visibility: jsii.String("visibility"),
//   												},
//   												Expression: jsii.String("expression"),
//   											},
//   										},
//   										SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   											Color: jsii.String("color"),
//   											Visibility: jsii.String("visibility"),
//   										},
//   										Visibility: jsii.String("visibility"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					SectionId: jsii.String("sectionId"),
//
//   					// the properties below are optional
//   					PageBreakConfiguration: &SectionPageBreakConfigurationProperty{
//   						After: &SectionAfterPageBreakProperty{
//   							Status: jsii.String("status"),
//   						},
//   					},
//   					RepeatConfiguration: &BodySectionRepeatConfigurationProperty{
//   						DimensionConfigurations: []interface{}{
//   							&BodySectionRepeatDimensionConfigurationProperty{
//   								DynamicCategoryDimensionConfiguration: &BodySectionDynamicCategoryDimensionConfigurationProperty{
//   									Column: &ColumnIdentifierProperty{
//   										ColumnName: jsii.String("columnName"),
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									},
//
//   									// the properties below are optional
//   									Limit: jsii.Number(123),
//   									SortByMetrics: []interface{}{
//   										&ColumnSortProperty{
//   											Direction: jsii.String("direction"),
//   											SortBy: &ColumnIdentifierProperty{
//   												ColumnName: jsii.String("columnName"),
//   												DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   											},
//
//   											// the properties below are optional
//   											AggregationFunction: &AggregationFunctionProperty{
//   												AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   													SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   													ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   												},
//   												CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   												DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   												NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   													PercentileAggregation: &PercentileAggregationProperty{
//   														PercentileValue: jsii.Number(123),
//   													},
//   													SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   												},
//   											},
//   										},
//   									},
//   								},
//   								DynamicNumericDimensionConfiguration: &BodySectionDynamicNumericDimensionConfigurationProperty{
//   									Column: &ColumnIdentifierProperty{
//   										ColumnName: jsii.String("columnName"),
//   										DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   									},
//
//   									// the properties below are optional
//   									Limit: jsii.Number(123),
//   									SortByMetrics: []interface{}{
//   										&ColumnSortProperty{
//   											Direction: jsii.String("direction"),
//   											SortBy: &ColumnIdentifierProperty{
//   												ColumnName: jsii.String("columnName"),
//   												DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   											},
//
//   											// the properties below are optional
//   											AggregationFunction: &AggregationFunctionProperty{
//   												AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   													SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   													ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   												},
//   												CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   												DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   												NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   													PercentileAggregation: &PercentileAggregationProperty{
//   														PercentileValue: jsii.Number(123),
//   													},
//   													SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   												},
//   											},
//   										},
//   									},
//   								},
//   							},
//   						},
//   						NonRepeatingVisuals: []*string{
//   							jsii.String("nonRepeatingVisuals"),
//   						},
//   						PageBreakConfiguration: &BodySectionRepeatPageBreakConfigurationProperty{
//   							After: &SectionAfterPageBreakProperty{
//   								Status: jsii.String("status"),
//   							},
//   						},
//   					},
//   					Style: &SectionStyleProperty{
//   						Height: jsii.String("height"),
//   						Padding: &SpacingProperty{
//   							Bottom: jsii.String("bottom"),
//   							Left: jsii.String("left"),
//   							Right: jsii.String("right"),
//   							Top: jsii.String("top"),
//   						},
//   					},
//   				},
//   			},
//   			CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   				PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   					PaperMargin: &SpacingProperty{
//   						Bottom: jsii.String("bottom"),
//   						Left: jsii.String("left"),
//   						Right: jsii.String("right"),
//   						Top: jsii.String("top"),
//   					},
//   					PaperOrientation: jsii.String("paperOrientation"),
//   					PaperSize: jsii.String("paperSize"),
//   				},
//   			},
//   			FooterSections: []interface{}{
//   				&HeaderFooterSectionConfigurationProperty{
//   					Layout: &SectionLayoutConfigurationProperty{
//   						FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   							Elements: []interface{}{
//   								&FreeFormLayoutElementProperty{
//   									ElementId: jsii.String("elementId"),
//   									ElementType: jsii.String("elementType"),
//   									Height: jsii.String("height"),
//   									Width: jsii.String("width"),
//   									XAxisLocation: jsii.String("xAxisLocation"),
//   									YAxisLocation: jsii.String("yAxisLocation"),
//
//   									// the properties below are optional
//   									BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									LoadingAnimation: &LoadingAnimationProperty{
//   										Visibility: jsii.String("visibility"),
//   									},
//   									RenderingRules: []interface{}{
//   										&SheetElementRenderingRuleProperty{
//   											ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   												Visibility: jsii.String("visibility"),
//   											},
//   											Expression: jsii.String("expression"),
//   										},
//   									},
//   									SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   						},
//   					},
//   					SectionId: jsii.String("sectionId"),
//
//   					// the properties below are optional
//   					Style: &SectionStyleProperty{
//   						Height: jsii.String("height"),
//   						Padding: &SpacingProperty{
//   							Bottom: jsii.String("bottom"),
//   							Left: jsii.String("left"),
//   							Right: jsii.String("right"),
//   							Top: jsii.String("top"),
//   						},
//   					},
//   				},
//   			},
//   			HeaderSections: []interface{}{
//   				&HeaderFooterSectionConfigurationProperty{
//   					Layout: &SectionLayoutConfigurationProperty{
//   						FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   							Elements: []interface{}{
//   								&FreeFormLayoutElementProperty{
//   									ElementId: jsii.String("elementId"),
//   									ElementType: jsii.String("elementType"),
//   									Height: jsii.String("height"),
//   									Width: jsii.String("width"),
//   									XAxisLocation: jsii.String("xAxisLocation"),
//   									YAxisLocation: jsii.String("yAxisLocation"),
//
//   									// the properties below are optional
//   									BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									LoadingAnimation: &LoadingAnimationProperty{
//   										Visibility: jsii.String("visibility"),
//   									},
//   									RenderingRules: []interface{}{
//   										&SheetElementRenderingRuleProperty{
//   											ConfigurationOverrides: &SheetElementConfigurationOverridesProperty{
//   												Visibility: jsii.String("visibility"),
//   											},
//   											Expression: jsii.String("expression"),
//   										},
//   									},
//   									SelectedBorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									Visibility: jsii.String("visibility"),
//   								},
//   							},
//   						},
//   					},
//   					SectionId: jsii.String("sectionId"),
//
//   					// the properties below are optional
//   					Style: &SectionStyleProperty{
//   						Height: jsii.String("height"),
//   						Padding: &SpacingProperty{
//   							Bottom: jsii.String("bottom"),
//   							Left: jsii.String("left"),
//   							Right: jsii.String("right"),
//   							Top: jsii.String("top"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layout.html
//
type CfnTemplate_LayoutProperty struct {
	// The configuration that determines what the type of layout for a sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-layout.html#cfn-quicksight-template-layout-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

