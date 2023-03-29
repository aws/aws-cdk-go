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
type CfnDashboard_DataPathLabelTypeProperty struct {
	// `CfnDashboard.DataPathLabelTypeProperty.FieldId`.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// `CfnDashboard.DataPathLabelTypeProperty.FieldValue`.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// `CfnDashboard.DataPathLabelTypeProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

