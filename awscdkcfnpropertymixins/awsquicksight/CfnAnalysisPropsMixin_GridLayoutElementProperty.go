package awsquicksight


// An element within a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gridLayoutElementProperty := &GridLayoutElementProperty{
//   	BackgroundStyle: &GridLayoutElementBackgroundStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	BorderRadius: jsii.String("borderRadius"),
//   	BorderStyle: &GridLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   		Width: jsii.String("width"),
//   	},
//   	ColumnIndex: jsii.Number(123),
//   	ColumnSpan: jsii.Number(123),
//   	ElementId: jsii.String("elementId"),
//   	ElementType: jsii.String("elementType"),
//   	LoadingAnimation: &LoadingAnimationProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Padding: jsii.String("padding"),
//   	RowIndex: jsii.Number(123),
//   	RowSpan: jsii.Number(123),
//   	SelectedBorderStyle: &GridLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   		Width: jsii.String("width"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html
//
type CfnAnalysisPropsMixin_GridLayoutElementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-backgroundstyle
	//
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-borderradius
	//
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-borderstyle
	//
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// The column index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-columnindex
	//
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// The width of a grid element expressed as a number of grid columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-columnspan
	//
	ColumnSpan *float64 `field:"optional" json:"columnSpan" yaml:"columnSpan"`
	// A unique identifier for an element within a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-elementid
	//
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// The type of element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-elementtype
	//
	ElementType *string `field:"optional" json:"elementType" yaml:"elementType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-loadinganimation
	//
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-padding
	//
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
	// The row index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-rowindex
	//
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
	// The height of a grid element expressed as a number of grid rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-rowspan
	//
	RowSpan *float64 `field:"optional" json:"rowSpan" yaml:"rowSpan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutelement.html#cfn-quicksight-analysis-gridlayoutelement-selectedborderstyle
	//
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
}

