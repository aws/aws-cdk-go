package previewawssecurityhubmixins


// Properties for CfnConfigurationPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationPolicyMixinProps := &CfnConfigurationPolicyMixinProps{
//   	ConfigurationPolicy: &PolicyProperty{
//   		SecurityHub: &SecurityHubPolicyProperty{
//   			EnabledStandardIdentifiers: []*string{
//   				jsii.String("enabledStandardIdentifiers"),
//   			},
//   			SecurityControlsConfiguration: &SecurityControlsConfigurationProperty{
//   				DisabledSecurityControlIdentifiers: []*string{
//   					jsii.String("disabledSecurityControlIdentifiers"),
//   				},
//   				EnabledSecurityControlIdentifiers: []*string{
//   					jsii.String("enabledSecurityControlIdentifiers"),
//   				},
//   				SecurityControlCustomParameters: []interface{}{
//   					&SecurityControlCustomParameterProperty{
//   						Parameters: map[string]interface{}{
//   							"parametersKey": &ParameterConfigurationProperty{
//   								"value": &ParameterValueProperty{
//   									"boolean": jsii.Boolean(false),
//   									"double": jsii.Number(123),
//   									"enum": jsii.String("enum"),
//   									"enumList": []*string{
//   										jsii.String("enumList"),
//   									},
//   									"integer": jsii.Number(123),
//   									"integerList": []interface{}{
//   										jsii.Number(123),
//   									},
//   									"string": jsii.String("string"),
//   									"stringList": []*string{
//   										jsii.String("stringList"),
//   									},
//   								},
//   								"valueType": jsii.String("valueType"),
//   							},
//   						},
//   						SecurityControlId: jsii.String("securityControlId"),
//   					},
//   				},
//   			},
//   			ServiceEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html
//
type CfnConfigurationPolicyMixinProps struct {
	// An object that defines how Security Hub is configured.
	//
	// It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-configurationpolicy
	//
	ConfigurationPolicy interface{} `field:"optional" json:"configurationPolicy" yaml:"configurationPolicy"`
	// The description of the configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration policy.
	//
	// Alphanumeric characters and the following ASCII characters are permitted: `-, ., !, *, /` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// User-defined tags associated with a configuration policy.
	//
	// For more information, see [Tagging Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

