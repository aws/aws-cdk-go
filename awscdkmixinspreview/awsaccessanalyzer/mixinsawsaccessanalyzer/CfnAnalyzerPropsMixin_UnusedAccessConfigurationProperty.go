package mixinsawsaccessanalyzer


// Contains information about an unused access analyzer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   unusedAccessConfigurationProperty := &UnusedAccessConfigurationProperty{
//   	AnalysisRule: &AnalysisRuleProperty{
//   		Exclusions: []interface{}{
//   			&AnalysisRuleCriteriaProperty{
//   				AccountIds: []*string{
//   					jsii.String("accountIds"),
//   				},
//   				ResourceTags: []interface{}{
//   					[]interface{}{
//   						&CfnTag{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	UnusedAccessAge: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html
//
type CfnAnalyzerPropsMixin_UnusedAccessConfigurationProperty struct {
	// Contains information about analysis rules for the analyzer.
	//
	// Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html#cfn-accessanalyzer-analyzer-unusedaccessconfiguration-analysisrule
	//
	AnalysisRule interface{} `field:"optional" json:"analysisRule" yaml:"analysisRule"`
	// The specified access age in days for which to generate findings for unused access.
	//
	// For example, if you specify 90 days, the analyzer will generate findings for IAM entities within the accounts of the selected organization for any access that hasn't been used in 90 or more days since the analyzer's last scan. You can choose a value between 1 and 365 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-unusedaccessconfiguration.html#cfn-accessanalyzer-analyzer-unusedaccessconfiguration-unusedaccessage
	//
	UnusedAccessAge *float64 `field:"optional" json:"unusedAccessAge" yaml:"unusedAccessAge"`
}

