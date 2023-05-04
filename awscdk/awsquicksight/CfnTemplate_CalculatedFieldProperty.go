package awsquicksight


// The calculated field of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedFieldProperty := &CalculatedFieldProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	Expression: jsii.String("expression"),
//   	Name: jsii.String("name"),
//   }
//
type CfnTemplate_CalculatedFieldProperty struct {
	// The data set that is used in this calculated field.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The expression of the calculated field.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the calculated field.
	Name *string `field:"required" json:"name" yaml:"name"`
}

