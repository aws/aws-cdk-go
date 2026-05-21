package awsconfig


// The value is either a dynamic (resource) value or a static value.
//
// You must select either a dynamic value or a static value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remediationParameterValueProperty := &RemediationParameterValueProperty{
//   	ResourceValue: &ResourceValueProperty{
//   		Value: jsii.String("value"),
//   	},
//   	StaticValue: &StaticValueProperty{
//   		Value: []*string{
//   			jsii.String("value"),
//   		},
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html
//
type CfnRemediationConfiguration_RemediationParameterValueProperty struct {
	// The value is dynamic and changes at run-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html#cfn-config-remediationconfiguration-remediationparametervalue-resourcevalue
	//
	ResourceValue interface{} `field:"optional" json:"resourceValue" yaml:"resourceValue"`
	// The value is static and does not change at run-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-remediationparametervalue.html#cfn-config-remediationconfiguration-remediationparametervalue-staticvalue
	//
	StaticValue interface{} `field:"optional" json:"staticValue" yaml:"staticValue"`
}

