package previewawsquicksightmixins


// The subtotal options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subtotalOptionsProperty := &SubtotalOptionsProperty{
//   	CustomLabel: jsii.String("customLabel"),
//   	FieldLevel: jsii.String("fieldLevel"),
//   	FieldLevelOptions: []interface{}{
//   		&PivotTableFieldSubtotalOptionsProperty{
//   			FieldId: jsii.String("fieldId"),
//   		},
//   	},
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
//   	StyleTargets: []interface{}{
//   		&TableStyleTargetProperty{
//   			CellType: jsii.String("cellType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html
//
type CfnAnalysisPropsMixin_SubtotalOptionsProperty struct {
	// The custom label string for the subtotal cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The field level (all, custom, last) for the subtotal cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-fieldlevel
	//
	FieldLevel *string `field:"optional" json:"fieldLevel" yaml:"fieldLevel"`
	// The optional configuration of subtotal cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-fieldleveloptions
	//
	FieldLevelOptions interface{} `field:"optional" json:"fieldLevelOptions" yaml:"fieldLevelOptions"`
	// The cell styling options for the subtotals of header cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-metricheadercellstyle
	//
	MetricHeaderCellStyle interface{} `field:"optional" json:"metricHeaderCellStyle" yaml:"metricHeaderCellStyle"`
	// The style targets options for subtotals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-styletargets
	//
	StyleTargets interface{} `field:"optional" json:"styleTargets" yaml:"styleTargets"`
	// The cell styling options for the subtotal cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-totalcellstyle
	//
	TotalCellStyle interface{} `field:"optional" json:"totalCellStyle" yaml:"totalCellStyle"`
	// The visibility configuration for the subtotal cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-totalsvisibility
	//
	TotalsVisibility *string `field:"optional" json:"totalsVisibility" yaml:"totalsVisibility"`
	// The cell styling options for the subtotals of value cells.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-subtotaloptions.html#cfn-quicksight-analysis-subtotaloptions-valuecellstyle
	//
	ValueCellStyle interface{} `field:"optional" json:"valueCellStyle" yaml:"valueCellStyle"`
}

