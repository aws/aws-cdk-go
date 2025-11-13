package awssecurityhub


// An object that defines how Security Hub is configured.
//
// It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyProperty := &PolicyProperty{
//   	SecurityHub: &SecurityHubPolicyProperty{
//   		EnabledStandardIdentifiers: []*string{
//   			jsii.String("enabledStandardIdentifiers"),
//   		},
//   		SecurityControlsConfiguration: &SecurityControlsConfigurationProperty{
//   			DisabledSecurityControlIdentifiers: []*string{
//   				jsii.String("disabledSecurityControlIdentifiers"),
//   			},
//   			EnabledSecurityControlIdentifiers: []*string{
//   				jsii.String("enabledSecurityControlIdentifiers"),
//   			},
//   			SecurityControlCustomParameters: []interface{}{
//   				&SecurityControlCustomParameterProperty{
//   					Parameters: map[string]interface{}{
//   						"parametersKey": &ParameterConfigurationProperty{
//   							"valueType": jsii.String("valueType"),
//
//   							// the properties below are optional
//   							"value": &ParameterValueProperty{
//   								"boolean": jsii.Boolean(false),
//   								"double": jsii.Number(123),
//   								"enum": jsii.String("enum"),
//   								"enumList": []*string{
//   									jsii.String("enumList"),
//   								},
//   								"integer": jsii.Number(123),
//   								"integerList": []interface{}{
//   									jsii.Number(123),
//   								},
//   								"string": jsii.String("string"),
//   								"stringList": []*string{
//   									jsii.String("stringList"),
//   								},
//   							},
//   						},
//   					},
//   					SecurityControlId: jsii.String("securityControlId"),
//   				},
//   			},
//   		},
//   		ServiceEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-policy.html
//
type CfnConfigurationPolicy_PolicyProperty struct {
	// The AWS service that the configuration policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-policy.html#cfn-securityhub-configurationpolicy-policy-securityhub
	//
	SecurityHub interface{} `field:"optional" json:"securityHub" yaml:"securityHub"`
}

