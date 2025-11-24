package mixinsawsquicksight


// The configuration that determines what the type of layout will be used on a sheet.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   layoutConfigurationProperty := &LayoutConfigurationProperty{
//   	FreeFormLayout: &FreeFormLayoutConfigurationProperty{
//   		CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   			ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   				OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   			},
//   		},
//   		Elements: []interface{}{
//   			&FreeFormLayoutElementProperty{
//   				BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   					Color: jsii.String("color"),
//   					Visibility: jsii.String("visibility"),
//   				},
//   				ElementId: jsii.String("elementId"),
//   				ElementType: jsii.String("elementType"),
//   				Height: jsii.String("height"),
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
//   				Width: jsii.String("width"),
//   				XAxisLocation: jsii.String("xAxisLocation"),
//   				YAxisLocation: jsii.String("yAxisLocation"),
//   			},
//   		},
//   	},
//   	GridLayout: &GridLayoutConfigurationProperty{
//   		CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   			ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   				OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   				ResizeOption: jsii.String("resizeOption"),
//   			},
//   		},
//   		Elements: []interface{}{
//   			&GridLayoutElementProperty{
//   				ColumnIndex: jsii.Number(123),
//   				ColumnSpan: jsii.Number(123),
//   				ElementId: jsii.String("elementId"),
//   				ElementType: jsii.String("elementType"),
//   				RowIndex: jsii.Number(123),
//   				RowSpan: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SectionBasedLayout: &SectionBasedLayoutConfigurationProperty{
//   		BodySections: []interface{}{
//   			&BodySectionConfigurationProperty{
//   				Content: &BodySectionContentProperty{
//   					Layout: &SectionLayoutConfigurationProperty{
//   						FreeFormLayout: &FreeFormSectionLayoutConfigurationProperty{
//   							Elements: []interface{}{
//   								&FreeFormLayoutElementProperty{
//   									BackgroundStyle: &FreeFormLayoutElementBackgroundStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									BorderStyle: &FreeFormLayoutElementBorderStyleProperty{
//   										Color: jsii.String("color"),
//   										Visibility: jsii.String("visibility"),
//   									},
//   									ElementId: jsii.String("elementId"),
//   									ElementType: jsii.String("elementType"),
//   									Height: jsii.String("height"),
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
//   									Width: jsii.String("width"),
//   									XAxisLocation: jsii.String("xAxisLocation"),
//   									YAxisLocation: jsii.String("yAxisLocation"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				PageBreakConfiguration: &SectionPageBreakConfigurationProperty{
//   					After: &SectionAfterPageBreakProperty{
//   						Status: jsii.String("status"),
//   					},
//   				},
//   				RepeatConfiguration: &BodySectionRepeatConfigurationProperty{
//   					DimensionConfigurations: []interface{}{
//   						&BodySectionRepeatDimensionConfigurationProperty{
//   							DynamicCategoryDimensionConfiguration: &BodySectionDynamicCategoryDimensionConfigurationProperty{
//   								Column: &ColumnIdentifierProperty{
//   									ColumnName: jsii.String("columnName"),
//   									DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   								},
//   								Limit: jsii.Number(123),
//   								SortByMetrics: []interface{}{
//   									&ColumnSortProperty{
//   										AggregationFunction: &AggregationFunctionProperty{
//   											AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   												SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   												ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   											},
//   											CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   											DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   											NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   												PercentileAggregation: &PercentileAggregationProperty{
//   													PercentileValue: jsii.Number(123),
//   												},
//   												SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   											},
//   										},
//   										Direction: jsii.String("direction"),
//   										SortBy: &ColumnIdentifierProperty{
//   											ColumnName: jsii.String("columnName"),
//   											DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   										},
//   									},
//   								},
//   							},
//   							DynamicNumericDimensionConfiguration: &BodySectionDynamicNumericDimensionConfigurationProperty{
//   								Column: &ColumnIdentifierProperty{
//   									ColumnName: jsii.String("columnName"),
//   									DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   								},
//   								Limit: jsii.Number(123),
//   								SortByMetrics: []interface{}{
//   									&ColumnSortProperty{
//   										AggregationFunction: &AggregationFunctionProperty{
//   											AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   												SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   												ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   											},
//   											CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   											DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   											NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   												PercentileAggregation: &PercentileAggregationProperty{
//   													PercentileValue: jsii.Number(123),
//   												},
//   												SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   											},
//   										},
//   										Direction: jsii.String("direction"),
//   										SortBy: &ColumnIdentifierProperty{
//   											ColumnName: jsii.String("columnName"),
//   											DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   					NonRepeatingVisuals: []*string{
//   						jsii.String("nonRepeatingVisuals"),
//   					},
//   					PageBreakConfiguration: &BodySectionRepeatPageBreakConfigurationProperty{
//   						After: &SectionAfterPageBreakProperty{
//   							Status: jsii.String("status"),
//   						},
//   					},
//   				},
//   				SectionId: jsii.String("sectionId"),
//   				Style: &SectionStyleProperty{
//   					Height: jsii.String("height"),
//   					Padding: &SpacingProperty{
//   						Bottom: jsii.String("bottom"),
//   						Left: jsii.String("left"),
//   						Right: jsii.String("right"),
//   						Top: jsii.String("top"),
//   					},
//   				},
//   			},
//   		},
//   		CanvasSizeOptions: &SectionBasedLayoutCanvasSizeOptionsProperty{
//   			PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   				PaperMargin: &SpacingProperty{
//   					Bottom: jsii.String("bottom"),
//   					Left: jsii.String("left"),
//   					Right: jsii.String("right"),
//   					Top: jsii.String("top"),
//   				},
//   				PaperOrientation: jsii.String("paperOrientation"),
//   				PaperSize: jsii.String("paperSize"),
//   			},
//   		},
//   		FooterSections: []interface{}{
//   			&HeaderFooterSectionConfigurationProperty{
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
//   				SectionId: jsii.String("sectionId"),
//   				Style: &SectionStyleProperty{
//   					Height: jsii.String("height"),
//   					Padding: &SpacingProperty{
//   						Bottom: jsii.String("bottom"),
//   						Left: jsii.String("left"),
//   						Right: jsii.String("right"),
//   						Top: jsii.String("top"),
//   					},
//   				},
//   			},
//   		},
//   		HeaderSections: []interface{}{
//   			&HeaderFooterSectionConfigurationProperty{
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
//   				SectionId: jsii.String("sectionId"),
//   				Style: &SectionStyleProperty{
//   					Height: jsii.String("height"),
//   					Padding: &SpacingProperty{
//   						Bottom: jsii.String("bottom"),
//   						Left: jsii.String("left"),
//   						Right: jsii.String("right"),
//   						Top: jsii.String("top"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html
//
type CfnDashboardPropsMixin_LayoutConfigurationProperty struct {
	// A free-form is optimized for a fixed width and has more control over the exact placement of layout elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-freeformlayout
	//
	FreeFormLayout interface{} `field:"optional" json:"freeFormLayout" yaml:"freeFormLayout"`
	// A type of layout that can be used on a sheet.
	//
	// In a grid layout, visuals snap to a grid with standard spacing and alignment. Dashboards are displayed as designed, with options to fit to screen or view at actual size. A grid layout can be configured to behave in one of two ways when the viewport is resized: `FIXED` or `RESPONSIVE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-gridlayout
	//
	GridLayout interface{} `field:"optional" json:"gridLayout" yaml:"gridLayout"`
	// A section based layout organizes visuals into multiple sections and has customized header, footer and page break.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-layoutconfiguration.html#cfn-quicksight-dashboard-layoutconfiguration-sectionbasedlayout
	//
	SectionBasedLayout interface{} `field:"optional" json:"sectionBasedLayout" yaml:"sectionBasedLayout"`
}

