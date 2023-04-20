package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingDateConfigurationProperty := &RollingDateConfigurationProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   }
//
type CfnDashboard_RollingDateConfigurationProperty struct {
	// `CfnDashboard.RollingDateConfigurationProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// `CfnDashboard.RollingDateConfigurationProperty.DataSetIdentifier`.
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}

