package awsaccessanalyzer


// The configuration for the analyzer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyzerConfigurationProperty := &AnalyzerConfigurationProperty{
//   	UnusedAccessConfiguration: &UnusedAccessConfigurationProperty{
//   		UnusedAccessAge: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html
//
type CfnAnalyzer_AnalyzerConfigurationProperty struct {
	// The Configuration for Unused Access Analyzer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html#cfn-accessanalyzer-analyzer-analyzerconfiguration-unusedaccessconfiguration
	//
	UnusedAccessConfiguration interface{} `field:"optional" json:"unusedAccessConfiguration" yaml:"unusedAccessConfiguration"`
}

