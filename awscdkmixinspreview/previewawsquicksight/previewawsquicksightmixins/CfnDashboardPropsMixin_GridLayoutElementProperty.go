package previewawsquicksightmixins


// An element within a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html
//
type CfnDashboardPropsMixin_GridLayoutElementProperty struct {
	// The background style configuration of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-backgroundstyle
	//
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// The border radius of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-borderradius
	//
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// The border style configuration of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-borderstyle
	//
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// The column index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-columnindex
	//
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// The width of a grid element expressed as a number of grid columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-columnspan
	//
	ColumnSpan *float64 `field:"optional" json:"columnSpan" yaml:"columnSpan"`
	// A unique identifier for an element within a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-elementid
	//
	ElementId *string `field:"optional" json:"elementId" yaml:"elementId"`
	// The type of element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-elementtype
	//
	ElementType *string `field:"optional" json:"elementType" yaml:"elementType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-loadinganimation
	//
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// The padding of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-padding
	//
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
	// The row index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-rowindex
	//
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
	// The height of a grid element expressed as a number of grid rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-rowspan
	//
	RowSpan *float64 `field:"optional" json:"rowSpan" yaml:"rowSpan"`
	// The border style configuration of a grid layout element.
	//
	// This border style is used when the element is selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-selectedborderstyle
	//
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
}

