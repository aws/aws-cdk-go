package awsaccessanalyzer


// Contains information about the configuration of an analyzer for an AWS organization or account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyzerConfigurationProperty := &AnalyzerConfigurationProperty{
//   	UnusedAccessConfiguration: &UnusedAccessConfigurationProperty{
//   		AnalysisRule: &AnalysisRuleProperty{
//   			Exclusions: []interface{}{
//   				&AnalysisRuleCriteriaProperty{
//   					AccountIds: []*string{
//   						jsii.String("accountIds"),
//   					},
//   					ResourceTags: []interface{}{
//   						[]interface{}{
//   							&CfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		UnusedAccessAge: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html
//
type CfnAnalyzer_AnalyzerConfigurationProperty struct {
	// Specifies the configuration of an unused access analyzer for an AWS organization or account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html#cfn-accessanalyzer-analyzer-analyzerconfiguration-unusedaccessconfiguration
	//
	UnusedAccessConfiguration interface{} `field:"optional" json:"unusedAccessConfiguration" yaml:"unusedAccessConfiguration"`
}

