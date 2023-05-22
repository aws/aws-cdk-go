package awsquicksight


// The data path that needs to be sorted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathValueProperty := &DataPathValueProperty{
//   	FieldId: jsii.String("fieldId"),
//   	FieldValue: jsii.String("fieldValue"),
//   }
//
type CfnAnalysis_DataPathValueProperty struct {
	// The field ID of the field that needs to be sorted.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The actual value of the field that needs to be sorted.
	FieldValue *string `field:"required" json:"fieldValue" yaml:"fieldValue"`
}

