package mixinsawssecurityhub


// An object that defines which security controls are enabled in an Security Hub configuration policy.
//
// The enablement status of a control is aligned across all of the enabled standards in an account.
//
// This property is required only if `ServiceEnabled` is set to `true` in your configuration policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   securityControlsConfigurationProperty := &SecurityControlsConfigurationProperty{
//   	DisabledSecurityControlIdentifiers: []*string{
//   		jsii.String("disabledSecurityControlIdentifiers"),
//   	},
//   	EnabledSecurityControlIdentifiers: []*string{
//   		jsii.String("enabledSecurityControlIdentifiers"),
//   	},
//   	SecurityControlCustomParameters: []interface{}{
//   		&SecurityControlCustomParameterProperty{
//   			Parameters: map[string]interface{}{
//   				"parametersKey": &ParameterConfigurationProperty{
//   					"value": &ParameterValueProperty{
//   						"boolean": jsii.Boolean(false),
//   						"double": jsii.Number(123),
//   						"enum": jsii.String("enum"),
//   						"enumList": []*string{
//   							jsii.String("enumList"),
//   						},
//   						"integer": jsii.Number(123),
//   						"integerList": []interface{}{
//   							jsii.Number(123),
//   						},
//   						"string": jsii.String("string"),
//   						"stringList": []*string{
//   							jsii.String("stringList"),
//   						},
//   					},
//   					"valueType": jsii.String("valueType"),
//   				},
//   			},
//   			SecurityControlId: jsii.String("securityControlId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolsconfiguration.html
//
type CfnConfigurationPolicyPropsMixin_SecurityControlsConfigurationProperty struct {
	// A list of security controls that are disabled in the configuration policy.
	//
	// Provide only one of `EnabledSecurityControlIdentifiers` or `DisabledSecurityControlIdentifiers` .
	//
	// If you provide `DisabledSecurityControlIdentifiers` , Security Hub enables all other controls not in the list, and enables [AutoEnableControls](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_UpdateSecurityHubConfiguration.html#securityhub-UpdateSecurityHubConfiguration-request-AutoEnableControls) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolsconfiguration.html#cfn-securityhub-configurationpolicy-securitycontrolsconfiguration-disabledsecuritycontrolidentifiers
	//
	DisabledSecurityControlIdentifiers *[]*string `field:"optional" json:"disabledSecurityControlIdentifiers" yaml:"disabledSecurityControlIdentifiers"`
	// A list of security controls that are enabled in the configuration policy.
	//
	// Provide only one of `EnabledSecurityControlIdentifiers` or `DisabledSecurityControlIdentifiers` .
	//
	// If you provide `EnabledSecurityControlIdentifiers` , Security Hub disables all other controls not in the list, and disables [AutoEnableControls](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_UpdateSecurityHubConfiguration.html#securityhub-UpdateSecurityHubConfiguration-request-AutoEnableControls) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolsconfiguration.html#cfn-securityhub-configurationpolicy-securitycontrolsconfiguration-enabledsecuritycontrolidentifiers
	//
	EnabledSecurityControlIdentifiers *[]*string `field:"optional" json:"enabledSecurityControlIdentifiers" yaml:"enabledSecurityControlIdentifiers"`
	// A list of security controls and control parameter values that are included in a configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-configurationpolicy-securitycontrolsconfiguration.html#cfn-securityhub-configurationpolicy-securitycontrolsconfiguration-securitycontrolcustomparameters
	//
	SecurityControlCustomParameters interface{} `field:"optional" json:"securityControlCustomParameters" yaml:"securityControlCustomParameters"`
}

