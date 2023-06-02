package awsquicksight


// The selected field options for the pivot table field options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableFieldOptionProperty := &PivotTableFieldOptionProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	CustomLabel: jsii.String("customLabel"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_PivotTableFieldOptionProperty struct {
	// The field ID of the pivot table field.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The custom label of the pivot table field.
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The visibility of the pivot table field.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

