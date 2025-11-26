package previewawsquicksightmixins


// The optional configuration of totals cells in a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTotalOptionsProperty := &PivotTotalOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	MetricHeaderCellStyle: &TableCellStyleProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		Border: &GlobalTableBorderOptionsProperty{
//   			SideSpecificBorder: &TableSideBorderOptionsProperty{
//   				Bottom: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerHorizontal: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerVertical: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Left: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Right: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Top: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			UniformBorder: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   		},
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Height: jsii.Number(123),
//   		HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   		TextWrap: jsii.String("textWrap"),
//   		VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Placement: jsii.String("placement"),
//   	ScrollStatus: jsii.String("scrollStatus"),
//   	TotalAggregationOptions: []interface{}{
//   		&TotalAggregationOptionProperty{
//   			FieldId: jsii.String("fieldId"),
//   			TotalAggregationFunction: &TotalAggregationFunctionProperty{
//   				SimpleTotalAggregationFunction: jsii.String("simpleTotalAggregationFunction"),
//   			},
//   		},
//   	},
//   	TotalCellStyle: &TableCellStyleProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		Border: &GlobalTableBorderOptionsProperty{
//   			SideSpecificBorder: &TableSideBorderOptionsProperty{
//   				Bottom: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerHorizontal: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerVertical: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Left: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Right: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Top: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			UniformBorder: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   		},
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Height: jsii.Number(123),
//   		HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   		TextWrap: jsii.String("textWrap"),
//   		VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TotalsVisibility: jsii.String("totalsVisibility"),
//   	ValueCellStyle: &TableCellStyleProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		Border: &GlobalTableBorderOptionsProperty{
//   			SideSpecificBorder: &TableSideBorderOptionsProperty{
//   				Bottom: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerHorizontal: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				InnerVertical: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Left: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Right: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   				Top: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			UniformBorder: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   		},
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Height: jsii.Number(123),
//   		HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   		TextWrap: jsii.String("textWrap"),
//   		VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html
//
type CfnTemplatePropsMixin_PivotTotalOptionsProperty struct {
	// The custom label string for the total cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The cell styling options for the total of header cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-metricheadercellstyle
	//
	MetricHeaderCellStyle interface{} `field:"optional" json:"metricHeaderCellStyle" yaml:"metricHeaderCellStyle"`
	// The placement (start, end) for the total cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-placement
	//
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
	// The scroll status (pinned, scrolled) for the total cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-scrollstatus
	//
	ScrollStatus *string `field:"optional" json:"scrollStatus" yaml:"scrollStatus"`
	// The total aggregation options for each value field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-totalaggregationoptions
	//
	TotalAggregationOptions interface{} `field:"optional" json:"totalAggregationOptions" yaml:"totalAggregationOptions"`
	// The cell styling options for the total cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-totalcellstyle
	//
	TotalCellStyle interface{} `field:"optional" json:"totalCellStyle" yaml:"totalCellStyle"`
	// The visibility configuration for the total cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-totalsvisibility
	//
	TotalsVisibility *string `field:"optional" json:"totalsVisibility" yaml:"totalsVisibility"`
	// The cell styling options for the totals of value cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottotaloptions.html#cfn-quicksight-template-pivottotaloptions-valuecellstyle
	//
	ValueCellStyle interface{} `field:"optional" json:"valueCellStyle" yaml:"valueCellStyle"`
}

