package awsinspectorv2


// Properties for defining a `CfnCodeSecurityScanConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeSecurityScanConfigurationProps := &CfnCodeSecurityScanConfigurationProps{
//   	Configuration: &CodeSecurityScanConfigurationProperty{
//   		RuleSetCategories: []*string{
//   			jsii.String("ruleSetCategories"),
//   		},
//
//   		// the properties below are optional
//   		ContinuousIntegrationScanConfiguration: &ContinuousIntegrationScanConfigurationProperty{
//   			SupportedEvents: []*string{
//   				jsii.String("supportedEvents"),
//   			},
//   		},
//   		PeriodicScanConfiguration: &PeriodicScanConfigurationProperty{
//   			Frequency: jsii.String("frequency"),
//   			FrequencyExpression: jsii.String("frequencyExpression"),
//   		},
//   	},
//   	Level: jsii.String("level"),
//   	Name: jsii.String("name"),
//   	ScopeSettings: &ScopeSettingsProperty{
//   		ProjectSelectionScope: jsii.String("projectSelectionScope"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html
//
type CfnCodeSecurityScanConfigurationProps struct {
	// The configuration settings for the code security scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The security level for the scan configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The name of the scan configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The scope settings that define which repositories will be scanned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-scopesettings
	//
	ScopeSettings interface{} `field:"optional" json:"scopeSettings" yaml:"scopeSettings"`
	// The tags to apply to the scan configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

