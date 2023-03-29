package awsquicksight


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
type CfnAnalysis_FieldSortProperty struct {
	// `CfnAnalysis.FieldSortProperty.Direction`.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// `CfnAnalysis.FieldSortProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

