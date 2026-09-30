package awsquicksight


// An element within a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutElementProperty := &GridLayoutElementProperty{
//   	ColumnSpan: jsii.Number(123),
//   	ElementId: jsii.String("elementId"),
//   	ElementType: jsii.String("elementType"),
//   	RowSpan: jsii.Number(123),
//
//   	// the properties below are optional
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
//   	LoadingAnimation: &LoadingAnimationProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Padding: jsii.String("padding"),
//   	RowIndex: jsii.Number(123),
//   	SelectedBorderStyle: &GridLayoutElementBorderStyleProperty{
//   		Color: jsii.String("color"),
//   		Visibility: jsii.String("visibility"),
//   		Width: jsii.String("width"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html
//
type CfnTemplate_GridLayoutElementProperty struct {
	// The width of a grid element expressed as a number of grid columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-columnspan
	//
	ColumnSpan *float64 `field:"required" json:"columnSpan" yaml:"columnSpan"`
	// A unique identifier for an element within a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-elementid
	//
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// The type of element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-elementtype
	//
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// The height of a grid element expressed as a number of grid rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-rowspan
	//
	RowSpan *float64 `field:"required" json:"rowSpan" yaml:"rowSpan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-backgroundstyle
	//
	BackgroundStyle interface{} `field:"optional" json:"backgroundStyle" yaml:"backgroundStyle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-borderradius
	//
	BorderRadius *string `field:"optional" json:"borderRadius" yaml:"borderRadius"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-borderstyle
	//
	BorderStyle interface{} `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// The column index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-columnindex
	//
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-loadinganimation
	//
	LoadingAnimation interface{} `field:"optional" json:"loadingAnimation" yaml:"loadingAnimation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-padding
	//
	Padding *string `field:"optional" json:"padding" yaml:"padding"`
	// The row index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-rowindex
	//
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelement.html#cfn-quicksight-template-gridlayoutelement-selectedborderstyle
	//
	SelectedBorderStyle interface{} `field:"optional" json:"selectedBorderStyle" yaml:"selectedBorderStyle"`
}

