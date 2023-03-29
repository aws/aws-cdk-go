package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathLabelTypeProperty := &DataPathLabelTypeProperty{
//   	FieldId: jsii.String("fieldId"),
//   	FieldValue: jsii.String("fieldValue"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_DataPathLabelTypeProperty struct {
	// `CfnAnalysis.DataPathLabelTypeProperty.FieldId`.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.DataPathLabelTypeProperty.FieldValue`.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// `CfnAnalysis.DataPathLabelTypeProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

