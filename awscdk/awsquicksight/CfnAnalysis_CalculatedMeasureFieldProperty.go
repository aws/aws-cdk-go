package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedMeasureFieldProperty := &CalculatedMeasureFieldProperty{
//   	Expression: jsii.String("expression"),
//   	FieldId: jsii.String("fieldId"),
//   }
//
type CfnAnalysis_CalculatedMeasureFieldProperty struct {
	// `CfnAnalysis.CalculatedMeasureFieldProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnAnalysis.CalculatedMeasureFieldProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

