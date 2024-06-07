package awssecurityhub


// Properties for defining a `CfnConfigurationPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationPolicyProps := &CfnConfigurationPolicyProps{
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
//   								"valueType": jsii.String("valueType"),
//
//   								// the properties below are optional
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
//   							},
//   						},
//   						SecurityControlId: jsii.String("securityControlId"),
//   					},
//   				},
//   			},
//   			ServiceEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html
//
type CfnConfigurationPolicyProps struct {
	// An object that defines how Security Hub is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-configurationpolicy
	//
	ConfigurationPolicy interface{} `field:"required" json:"configurationPolicy" yaml:"configurationPolicy"`
	// The name of the configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the configuration policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A key-value pair to associate with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-configurationpolicy.html#cfn-securityhub-configurationpolicy-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

