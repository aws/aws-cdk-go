package awsinspectorv2


// Contains the configuration settings for code security scans.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeSecurityScanConfigurationProperty := &CodeSecurityScanConfigurationProperty{
//   	RuleSetCategories: []*string{
//   		jsii.String("ruleSetCategories"),
//   	},
//
//   	// the properties below are optional
//   	ContinuousIntegrationScanConfiguration: &ContinuousIntegrationScanConfigurationProperty{
//   		SupportedEvents: []*string{
//   			jsii.String("supportedEvents"),
//   		},
//   	},
//   	PeriodicScanConfiguration: &PeriodicScanConfigurationProperty{
//   		Frequency: jsii.String("frequency"),
//   		FrequencyExpression: jsii.String("frequencyExpression"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html
//
type CfnCodeSecurityScanConfiguration_CodeSecurityScanConfigurationProperty struct {
	// The categories of security rules to be applied during the scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-rulesetcategories
	//
	RuleSetCategories *[]*string `field:"required" json:"ruleSetCategories" yaml:"ruleSetCategories"`
	// Configuration settings for continuous integration scans that run automatically when code changes are made.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-continuousintegrationscanconfiguration
	//
	ContinuousIntegrationScanConfiguration interface{} `field:"optional" json:"continuousIntegrationScanConfiguration" yaml:"continuousIntegrationScanConfiguration"`
	// Configuration settings for periodic scans that run on a scheduled basis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-periodicscanconfiguration
	//
	PeriodicScanConfiguration interface{} `field:"optional" json:"periodicScanConfiguration" yaml:"periodicScanConfiguration"`
}

