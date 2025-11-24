package mixinsawsquicksight


// The sort configuration for a field in a field well.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldSortProperty := &FieldSortProperty{
//   	Direction: jsii.String("direction"),
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fieldsort.html
//
type CfnDashboardPropsMixin_FieldSortProperty struct {
	// The sort direction. Choose one of the following options:.
	//
	// - `ASC` : Ascending
	// - `DESC` : Descending.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fieldsort.html#cfn-quicksight-dashboard-fieldsort-direction
	//
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The sort configuration target field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-fieldsort.html#cfn-quicksight-dashboard-fieldsort-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
}

