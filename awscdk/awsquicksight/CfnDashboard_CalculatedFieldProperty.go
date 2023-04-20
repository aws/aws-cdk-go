package awsquicksight


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
type CfnDashboard_CalculatedFieldProperty struct {
	// `CfnDashboard.CalculatedFieldProperty.DataSetIdentifier`.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// `CfnDashboard.CalculatedFieldProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnDashboard.CalculatedFieldProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

