package mixinsawsquicksight


// An element within a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gridLayoutElementProperty := &GridLayoutElementProperty{
//   	ColumnIndex: jsii.Number(123),
//   	ColumnSpan: jsii.Number(123),
//   	ElementId: jsii.String("elementId"),
//   	ElementType: jsii.String("elementType"),
//   	RowIndex: jsii.Number(123),
//   	RowSpan: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html
//
type CfnDashboardPropsMixin_GridLayoutElementProperty struct {
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
	// The row index for the upper left corner of an element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-rowindex
	//
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
	// The height of a grid element expressed as a number of grid rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelement.html#cfn-quicksight-dashboard-gridlayoutelement-rowspan
	//
	RowSpan *float64 `field:"optional" json:"rowSpan" yaml:"rowSpan"`
}

