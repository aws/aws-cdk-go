package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	Placement: jsii.String("placement"),
//   	ScrollStatus: jsii.String("scrollStatus"),
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
type CfnTemplate_PivotTotalOptionsProperty struct {
	// `CfnTemplate.PivotTotalOptionsProperty.CustomLabel`.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// `CfnTemplate.PivotTotalOptionsProperty.MetricHeaderCellStyle`.
	MetricHeaderCellStyle interface{} `field:"optional" json:"metricHeaderCellStyle" yaml:"metricHeaderCellStyle"`
	// `CfnTemplate.PivotTotalOptionsProperty.Placement`.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
	// `CfnTemplate.PivotTotalOptionsProperty.ScrollStatus`.
	ScrollStatus *string `field:"optional" json:"scrollStatus" yaml:"scrollStatus"`
	// `CfnTemplate.PivotTotalOptionsProperty.TotalCellStyle`.
	TotalCellStyle interface{} `field:"optional" json:"totalCellStyle" yaml:"totalCellStyle"`
	// `CfnTemplate.PivotTotalOptionsProperty.TotalsVisibility`.
	TotalsVisibility *string `field:"optional" json:"totalsVisibility" yaml:"totalsVisibility"`
	// `CfnTemplate.PivotTotalOptionsProperty.ValueCellStyle`.
	ValueCellStyle interface{} `field:"optional" json:"valueCellStyle" yaml:"valueCellStyle"`
}

