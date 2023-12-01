package awsaccessanalyzer


// The Configuration for Unused Access Analyzer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unusedAccessConfigurationProperty := &UnusedAccessConfigurationProperty{
//   	UnusedAccessAge: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html
//
type CfnAnalyzer_UnusedAccessConfigurationProperty struct {
	// The specified access age in days for which to generate findings for unused access.
	//
	// For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that haven't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 180 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html#cfn-accessanalyzer-analyzer-unusedaccessconfiguration-unusedaccessage
	//
	UnusedAccessAge *float64 `field:"optional" json:"unusedAccessAge" yaml:"unusedAccessAge"`
}

