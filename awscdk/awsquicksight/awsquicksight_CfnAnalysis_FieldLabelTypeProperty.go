package awsquicksight


// The field label type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldLabelTypeProperty := &FieldLabelTypeProperty{
//   	FieldId: jsii.String("fieldId"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_FieldLabelTypeProperty struct {
	// Indicates the field that is targeted by the field label.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The visibility of the field label.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

