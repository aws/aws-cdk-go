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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldsubtotaloptions.html
//
type CfnDashboard_PivotTableFieldSubtotalOptionsProperty struct {
	// The field ID of the subtotal options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldsubtotaloptions.html#cfn-quicksight-dashboard-pivottablefieldsubtotaloptions-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

