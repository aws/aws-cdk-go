package awsquicksight


// The sort configuration for a field in a field well.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldSortProperty := &FieldSortProperty{
//   	Direction: jsii.String("direction"),
//   	FieldId: jsii.String("fieldId"),
//   }
//
type CfnTemplate_FieldSortProperty struct {
	// The sort direction. Choose one of the following options:.
	//
	// - `ASC` : Ascending
	// - `DESC` : Descending.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The sort configuration target field.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

