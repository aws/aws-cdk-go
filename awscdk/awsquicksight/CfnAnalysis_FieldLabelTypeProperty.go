package awsquicksight


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
	// `CfnAnalysis.FieldLabelTypeProperty.FieldId`.
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.FieldLabelTypeProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

