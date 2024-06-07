package awssecurityhub


// An object that defines how AWS Security Hub is configured.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityHubPolicyProperty := &SecurityHubPolicyProperty{
//   	EnabledStandardIdentifiers: []*string{
//   		jsii.String("enabledStandardIdentifiers"),
//   	},
//   	SecurityControlsConfiguration: &SecurityControlsConfigurationProperty{
//   		DisabledSecurityControlIdentifiers: []*string{
//   			jsii.String("disabledSecurityControlIdentifiers"),
//   		},
//   		EnabledSecurityControlIdentifiers: []*string{
//   			jsii.String("enabledSecurityControlIdentifiers"),
//   		},
//   		SecurityControlCustomParameters: []interface{}{
//   			&SecurityControlCustomParameterProperty{
//   				Parameters: map[string]interface{}{
//   					"parametersKey": &ParameterConfigurationProperty{
//   						"valueType": jsii.String("valueType"),
//
//   						// the properties below are optional
//   						"value": &ParameterValueProperty{
//   							"boolean": jsii.Boolean(false),
//   							"double": jsii.Number(123),
//   							"enum": jsii.String("enum"),
//   							"enumList": []*string{
//   								jsii.String("enumList"),
//   							},
//   							"integer": jsii.Number(123),
//   							"integerList": []interface{}{
//   								jsii.Number(123),
//   							},
//   							"string": jsii.String("string"),
//   							"stringList": []*string{
//   								jsii.String("stringList"),
//   							},
//   						},
//   					},
//   				},
//   				SecurityControlId: jsii.String("securityControlId"),
//   			},
//   		},
//   	},
//   	ServiceEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html
//
type CfnConfigurationPolicy_SecurityHubPolicyProperty struct {
	// A list that defines which security standards are enabled in the configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-enabledstandardidentifiers
	//
	EnabledStandardIdentifiers *[]*string `field:"optional" json:"enabledStandardIdentifiers" yaml:"enabledStandardIdentifiers"`
	// An object that defines which security controls are enabled in an AWS Security Hub configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-securitycontrolsconfiguration
	//
	SecurityControlsConfiguration interface{} `field:"optional" json:"securityControlsConfiguration" yaml:"securityControlsConfiguration"`
	// Indicates whether Security Hub is enabled in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securityhubpolicy.html#cfn-securityhub-configurationpolicy-securityhubpolicy-serviceenabled
	//
	ServiceEnabled interface{} `field:"optional" json:"serviceEnabled" yaml:"serviceEnabled"`
}

