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
//   	ColumnIndex: jsii.Number(123),
//   	RowIndex: jsii.Number(123),
//   }
//
type CfnTemplate_GridLayoutElementProperty struct {
	// The width of a grid element expressed as a number of grid columns.
	ColumnSpan *float64 `field:"required" json:"columnSpan" yaml:"columnSpan"`
	// A unique identifier for an element within a grid layout.
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// The type of element.
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// The height of a grid element expressed as a number of grid rows.
	RowSpan *float64 `field:"required" json:"rowSpan" yaml:"rowSpan"`
	// The column index for the upper left corner of an element.
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// The row index for the upper left corner of an element.
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
}

