package awsquicksight


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
	// `CfnTemplate.GridLayoutElementProperty.ColumnSpan`.
	ColumnSpan *float64 `field:"required" json:"columnSpan" yaml:"columnSpan"`
	// `CfnTemplate.GridLayoutElementProperty.ElementId`.
	ElementId *string `field:"required" json:"elementId" yaml:"elementId"`
	// `CfnTemplate.GridLayoutElementProperty.ElementType`.
	ElementType *string `field:"required" json:"elementType" yaml:"elementType"`
	// `CfnTemplate.GridLayoutElementProperty.RowSpan`.
	RowSpan *float64 `field:"required" json:"rowSpan" yaml:"rowSpan"`
	// `CfnTemplate.GridLayoutElementProperty.ColumnIndex`.
	ColumnIndex *float64 `field:"optional" json:"columnIndex" yaml:"columnIndex"`
	// `CfnTemplate.GridLayoutElementProperty.RowIndex`.
	RowIndex *float64 `field:"optional" json:"rowIndex" yaml:"rowIndex"`
}

