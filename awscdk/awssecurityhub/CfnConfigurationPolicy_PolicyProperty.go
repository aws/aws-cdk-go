package awssecurityhub


// An object that defines how Security Hub is configured.
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
	// An object that defines how AWS Security Hub is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-policy.html#cfn-securityhub-configurationpolicy-policy-securityhub
	//
	SecurityHub interface{} `field:"optional" json:"securityHub" yaml:"securityHub"`
}

