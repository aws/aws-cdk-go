package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   			FontSize: &FontSizeProperty{
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
//   			FontSize: &FontSizeProperty{
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
//   			FontSize: &FontSizeProperty{
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
type CfnTemplate_SubtotalOptionsProperty struct {
	// `CfnTemplate.SubtotalOptionsProperty.CustomLabel`.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// `CfnTemplate.SubtotalOptionsProperty.FieldLevel`.
	FieldLevel *string `field:"optional" json:"fieldLevel" yaml:"fieldLevel"`
	// `CfnTemplate.SubtotalOptionsProperty.FieldLevelOptions`.
	FieldLevelOptions interface{} `field:"optional" json:"fieldLevelOptions" yaml:"fieldLevelOptions"`
	// `CfnTemplate.SubtotalOptionsProperty.MetricHeaderCellStyle`.
	MetricHeaderCellStyle interface{} `field:"optional" json:"metricHeaderCellStyle" yaml:"metricHeaderCellStyle"`
	// `CfnTemplate.SubtotalOptionsProperty.TotalCellStyle`.
	TotalCellStyle interface{} `field:"optional" json:"totalCellStyle" yaml:"totalCellStyle"`
	// `CfnTemplate.SubtotalOptionsProperty.TotalsVisibility`.
	TotalsVisibility *string `field:"optional" json:"totalsVisibility" yaml:"totalsVisibility"`
	// `CfnTemplate.SubtotalOptionsProperty.ValueCellStyle`.
	ValueCellStyle interface{} `field:"optional" json:"valueCellStyle" yaml:"valueCellStyle"`
}

