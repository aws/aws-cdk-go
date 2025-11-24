package mixinsawsinspectorv2


// Contains the configuration settings for code security scans.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeSecurityScanConfigurationProperty := &CodeSecurityScanConfigurationProperty{
//   	ContinuousIntegrationScanConfiguration: &ContinuousIntegrationScanConfigurationProperty{
//   		SupportedEvents: []*string{
//   			jsii.String("supportedEvents"),
//   		},
//   	},
//   	PeriodicScanConfiguration: &PeriodicScanConfigurationProperty{
//   		Frequency: jsii.String("frequency"),
//   		FrequencyExpression: jsii.String("frequencyExpression"),
//   	},
//   	RuleSetCategories: []*string{
//   		jsii.String("ruleSetCategories"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html
//
type CfnCodeSecurityScanConfigurationPropsMixin_CodeSecurityScanConfigurationProperty struct {
	// Configuration settings for continuous integration scans that run automatically when code changes are made.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-continuousintegrationscanconfiguration
	//
	ContinuousIntegrationScanConfiguration interface{} `field:"optional" json:"continuousIntegrationScanConfiguration" yaml:"continuousIntegrationScanConfiguration"`
	// Configuration settings for periodic scans that run on a scheduled basis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-periodicscanconfiguration
	//
	PeriodicScanConfiguration interface{} `field:"optional" json:"periodicScanConfiguration" yaml:"periodicScanConfiguration"`
	// The categories of security rules to be applied during the scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-rulesetcategories
	//
	RuleSetCategories *[]*string `field:"optional" json:"ruleSetCategories" yaml:"ruleSetCategories"`
}

