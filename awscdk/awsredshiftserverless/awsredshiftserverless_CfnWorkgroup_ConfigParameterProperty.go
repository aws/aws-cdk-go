package awsredshiftserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configParameterProperty := &configParameterProperty{
//   	parameterKey: jsii.String("parameterKey"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnWorkgroup_ConfigParameterProperty struct {
	// `CfnWorkgroup.ConfigParameterProperty.ParameterKey`.
	ParameterKey *string `field:"optional" json:"parameterKey" yaml:"parameterKey"`
	// `CfnWorkgroup.ConfigParameterProperty.ParameterValue`.
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

