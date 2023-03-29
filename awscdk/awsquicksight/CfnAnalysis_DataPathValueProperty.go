package awsquicksight


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
	// `CfnAnalysis.DataPathValueProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.DataPathValueProperty.FieldValue`.
	FieldValue *string `field:"required" json:"fieldValue" yaml:"fieldValue"`
}

