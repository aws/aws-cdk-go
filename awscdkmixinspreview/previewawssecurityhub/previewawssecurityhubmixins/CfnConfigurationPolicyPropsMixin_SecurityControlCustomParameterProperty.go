package previewawssecurityhubmixins


// A list of security controls and control parameter values that are included in a configuration policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityControlCustomParameterProperty := &SecurityControlCustomParameterProperty{
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterConfigurationProperty{
//   			"value": &ParameterValueProperty{
//   				"boolean": jsii.Boolean(false),
//   				"double": jsii.Number(123),
//   				"enum": jsii.String("enum"),
//   				"enumList": []*string{
//   					jsii.String("enumList"),
//   				},
//   				"integer": jsii.Number(123),
//   				"integerList": []interface{}{
//   					jsii.Number(123),
//   				},
//   				"string": jsii.String("string"),
//   				"stringList": []*string{
//   					jsii.String("stringList"),
//   				},
//   			},
//   			"valueType": jsii.String("valueType"),
//   		},
//   	},
//   	SecurityControlId: jsii.String("securityControlId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolcustomparameter.html
//
type CfnConfigurationPolicyPropsMixin_SecurityControlCustomParameterProperty struct {
	// An object that specifies parameter values for a control in a configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolcustomparameter.html#cfn-securityhub-configurationpolicy-securitycontrolcustomparameter-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The ID of the security control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolcustomparameter.html#cfn-securityhub-configurationpolicy-securitycontrolcustomparameter-securitycontrolid
	//
	SecurityControlId *string `field:"optional" json:"securityControlId" yaml:"securityControlId"`
}

