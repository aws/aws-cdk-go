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
type CfnTemplate_CalculatedMeasureFieldProperty struct {
	// `CfnTemplate.CalculatedMeasureFieldProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnTemplate.CalculatedMeasureFieldProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}

