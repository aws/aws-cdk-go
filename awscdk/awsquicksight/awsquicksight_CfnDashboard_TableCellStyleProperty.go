package awsquicksight


// The table cell style for a cell in pivot table or table visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableCellStyleProperty := &TableCellStyleProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	Border: &GlobalTableBorderOptionsProperty{
//   		SideSpecificBorder: &TableSideBorderOptionsProperty{
//   			Bottom: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   			InnerHorizontal: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   			InnerVertical: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   			Left: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   			Right: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   			Top: &TableBorderOptionsProperty{
//   				Color: jsii.String("color"),
//   				Style: jsii.String("style"),
//   				Thickness: jsii.Number(123),
//   			},
//   		},
//   		UniformBorder: &TableBorderOptionsProperty{
//   			Color: jsii.String("color"),
//   			Style: jsii.String("style"),
//   			Thickness: jsii.Number(123),
//   		},
//   	},
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Height: jsii.Number(123),
//   	HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   	TextWrap: jsii.String("textWrap"),
//   	VerticalTextAlignment: jsii.String("verticalTextAlignment"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_TableCellStyleProperty struct {
	// The background color for the table cells.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The borders for the table cells.
	Border interface{} `field:"optional" json:"border" yaml:"border"`
	// The font configuration of the table cells.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// The height color for the table cells.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The horizontal text alignment (left, center, right, auto) for the table cells.
	HorizontalTextAlignment *string `field:"optional" json:"horizontalTextAlignment" yaml:"horizontalTextAlignment"`
	// The text wrap (none, wrap) for the table cells.
	TextWrap *string `field:"optional" json:"textWrap" yaml:"textWrap"`
	// The vertical text alignment (top, middle, bottom) for the table cells.
	VerticalTextAlignment *string `field:"optional" json:"verticalTextAlignment" yaml:"verticalTextAlignment"`
	// The visibility of the table cells.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

