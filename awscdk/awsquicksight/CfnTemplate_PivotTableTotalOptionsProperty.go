package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableTotalOptionsProperty := &PivotTableTotalOptionsProperty{
//   	ColumnSubtotalOptions: &SubtotalOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FieldLevel: jsii.String("fieldLevel"),
//   		FieldLevelOptions: []interface{}{
//   			&PivotTableFieldSubtotalOptionsProperty{
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   		MetricHeaderCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalsVisibility: jsii.String("totalsVisibility"),
//   		ValueCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	ColumnTotalOptions: &PivotTotalOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		MetricHeaderCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		Placement: jsii.String("placement"),
//   		ScrollStatus: jsii.String("scrollStatus"),
//   		TotalCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalsVisibility: jsii.String("totalsVisibility"),
//   		ValueCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	RowSubtotalOptions: &SubtotalOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FieldLevel: jsii.String("fieldLevel"),
//   		FieldLevelOptions: []interface{}{
//   			&PivotTableFieldSubtotalOptionsProperty{
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   		MetricHeaderCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalsVisibility: jsii.String("totalsVisibility"),
//   		ValueCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	RowTotalOptions: &PivotTotalOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		MetricHeaderCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		Placement: jsii.String("placement"),
//   		ScrollStatus: jsii.String("scrollStatus"),
//   		TotalCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   		TotalsVisibility: jsii.String("totalsVisibility"),
//   		ValueCellStyle: &TableCellStyleProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			Border: &GlobalTableBorderOptionsProperty{
//   				SideSpecificBorder: &TableSideBorderOptionsProperty{
//   					Bottom: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerHorizontal: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					InnerVertical: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Left: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Right: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   					Top: &TableBorderOptionsProperty{
//   						Color: jsii.String("color"),
//   						Style: jsii.String("style"),
//   						Thickness: jsii.Number(123),
//   					},
//   				},
//   				UniformBorder: &TableBorderOptionsProperty{
//   					Color: jsii.String("color"),
//   					Style: jsii.String("style"),
//   					Thickness: jsii.Number(123),
//   				},
//   			},
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Height: jsii.Number(123),
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			TextWrap: jsii.String("textWrap"),
//   			VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   }
//
type CfnTemplate_PivotTableTotalOptionsProperty struct {
	// `CfnTemplate.PivotTableTotalOptionsProperty.ColumnSubtotalOptions`.
	ColumnSubtotalOptions interface{} `field:"optional" json:"columnSubtotalOptions" yaml:"columnSubtotalOptions"`
	// `CfnTemplate.PivotTableTotalOptionsProperty.ColumnTotalOptions`.
	ColumnTotalOptions interface{} `field:"optional" json:"columnTotalOptions" yaml:"columnTotalOptions"`
	// `CfnTemplate.PivotTableTotalOptionsProperty.RowSubtotalOptions`.
	RowSubtotalOptions interface{} `field:"optional" json:"rowSubtotalOptions" yaml:"rowSubtotalOptions"`
	// `CfnTemplate.PivotTableTotalOptionsProperty.RowTotalOptions`.
	RowTotalOptions interface{} `field:"optional" json:"rowTotalOptions" yaml:"rowTotalOptions"`
}

