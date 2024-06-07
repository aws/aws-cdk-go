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
//
//   	// the properties below are optional
//   	Value: &ParameterValueProperty{
//   		Boolean: jsii.Boolean(false),
//   		Double: jsii.Number(123),
//   		Enum: jsii.String("enum"),
//   		EnumList: []*string{
//   			jsii.String("enumList"),
//   		},
//   		Integer: jsii.Number(123),
//   		IntegerList: []interface{}{
//   			jsii.Number(123),
//   		},
//   		String: jsii.String("string"),
//   		StringList: []*string{
//   			jsii.String("stringList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html
//
type CfnConfigurationPolicy_ParameterConfigurationProperty struct {
	// Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html#cfn-securityhub-configurationpolicy-parameterconfiguration-valuetype
	//
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// An object that includes the data type of a security control parameter and its current value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-parameterconfiguration.html#cfn-securityhub-configurationpolicy-parameterconfiguration-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

