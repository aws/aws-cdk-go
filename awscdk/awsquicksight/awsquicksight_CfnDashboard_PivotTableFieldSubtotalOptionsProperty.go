package awsquicksight


// The optional configuration of subtotals cells.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldSubtotalOptionsProperty := &PivotTableFieldSubtotalOptionsProperty{
//   	FieldId: jsii.String("fieldId"),
//   }
//
type CfnDashboard_PivotTableFieldSubtotalOptionsProperty struct {
	// The field ID of the subtotal options.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

