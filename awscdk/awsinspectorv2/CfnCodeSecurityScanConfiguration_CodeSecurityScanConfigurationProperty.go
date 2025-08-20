package awsinspectorv2


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-rulesetcategories
	//
	RuleSetCategories *[]*string `field:"required" json:"ruleSetCategories" yaml:"ruleSetCategories"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-continuousintegrationscanconfiguration
	//
	ContinuousIntegrationScanConfiguration interface{} `field:"optional" json:"continuousIntegrationScanConfiguration" yaml:"continuousIntegrationScanConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-codesecurityscanconfiguration-periodicscanconfiguration
	//
	PeriodicScanConfiguration interface{} `field:"optional" json:"periodicScanConfiguration" yaml:"periodicScanConfiguration"`
}

