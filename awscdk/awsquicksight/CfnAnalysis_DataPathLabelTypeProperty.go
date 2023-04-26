package awsquicksight


// The option that specifies individual data values for labels.
//
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
	// The field ID of the field that the data label needs to be applied to.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The actual value of the field that is labeled.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// The visibility of the data label.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

