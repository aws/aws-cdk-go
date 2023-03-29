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
type CfnAnalysis_TableCellStyleProperty struct {
	// `CfnAnalysis.TableCellStyleProperty.BackgroundColor`.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// `CfnAnalysis.TableCellStyleProperty.Border`.
	Border interface{} `field:"optional" json:"border" yaml:"border"`
	// `CfnAnalysis.TableCellStyleProperty.FontConfiguration`.
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// `CfnAnalysis.TableCellStyleProperty.Height`.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// `CfnAnalysis.TableCellStyleProperty.HorizontalTextAlignment`.
	HorizontalTextAlignment *string `field:"optional" json:"horizontalTextAlignment" yaml:"horizontalTextAlignment"`
	// `CfnAnalysis.TableCellStyleProperty.TextWrap`.
	TextWrap *string `field:"optional" json:"textWrap" yaml:"textWrap"`
	// `CfnAnalysis.TableCellStyleProperty.VerticalTextAlignment`.
	VerticalTextAlignment *string `field:"optional" json:"verticalTextAlignment" yaml:"verticalTextAlignment"`
	// `CfnAnalysis.TableCellStyleProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

