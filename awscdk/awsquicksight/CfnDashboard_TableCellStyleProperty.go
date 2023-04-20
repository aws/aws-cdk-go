package awsquicksight


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
	// `CfnDashboard.TableCellStyleProperty.BackgroundColor`.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// `CfnDashboard.TableCellStyleProperty.Border`.
	Border interface{} `field:"optional" json:"border" yaml:"border"`
	// `CfnDashboard.TableCellStyleProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// `CfnDashboard.TableCellStyleProperty.Height`.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// `CfnDashboard.TableCellStyleProperty.HorizontalTextAlignment`.
	HorizontalTextAlignment *string `field:"optional" json:"horizontalTextAlignment" yaml:"horizontalTextAlignment"`
	// `CfnDashboard.TableCellStyleProperty.TextWrap`.
	TextWrap *string `field:"optional" json:"textWrap" yaml:"textWrap"`
	// `CfnDashboard.TableCellStyleProperty.VerticalTextAlignment`.
	VerticalTextAlignment *string `field:"optional" json:"verticalTextAlignment" yaml:"verticalTextAlignment"`
	// `CfnDashboard.TableCellStyleProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

