package awssecurityhub


// An object that provides the current value of a security control parameter and identifies whether it has been customized.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterConfigurationProperty := &ParameterConfigurationProperty{
//   	ValueType: jsii.String("valueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parameterconfiguration.html
//
type CfnSecurityControl_ParameterConfigurationProperty struct {
	// Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.
	//
	// When `ValueType` is set equal to `DEFAULT` , the default behavior can be a specific Security Hub default value, or the default behavior can be to ignore a specific parameter. When `ValueType` is set equal to `DEFAULT` , Security Hub ignores user-provided input for the `Value` field.
	//
	// When `ValueType` is set equal to `CUSTOM` , the `Value` field can't be empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parameterconfiguration.html#cfn-securityhub-securitycontrol-parameterconfiguration-valuetype
	//
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
}

