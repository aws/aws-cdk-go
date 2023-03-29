package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableOptionsProperty := &TableOptionsProperty{
//   	CellStyle: &TableCellStyleProperty{
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
//   	HeaderStyle: &TableCellStyleProperty{
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
//   	Orientation: jsii.String("orientation"),
//   	RowAlternateColorOptions: &RowAlternateColorOptionsProperty{
//   		RowAlternateColors: []*string{
//   			jsii.String("rowAlternateColors"),
//   		},
//   		Status: jsii.String("status"),
//   	},
//   }
//
type CfnDashboard_TableOptionsProperty struct {
	// `CfnDashboard.TableOptionsProperty.CellStyle`.
	CellStyle interface{} `field:"optional" json:"cellStyle" yaml:"cellStyle"`
	// `CfnDashboard.TableOptionsProperty.HeaderStyle`.
	HeaderStyle interface{} `field:"optional" json:"headerStyle" yaml:"headerStyle"`
	// `CfnDashboard.TableOptionsProperty.Orientation`.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// `CfnDashboard.TableOptionsProperty.RowAlternateColorOptions`.
	RowAlternateColorOptions interface{} `field:"optional" json:"rowAlternateColorOptions" yaml:"rowAlternateColorOptions"`
}

