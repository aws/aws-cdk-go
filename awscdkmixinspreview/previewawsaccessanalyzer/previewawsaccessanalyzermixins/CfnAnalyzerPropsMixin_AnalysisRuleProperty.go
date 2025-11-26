package previewawsaccessanalyzermixins


// Contains information about analysis rules for the analyzer.
//
// Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisRuleProperty := &AnalysisRuleProperty{
//   	Exclusions: []interface{}{
//   		&AnalysisRuleCriteriaProperty{
//   			AccountIds: []*string{
//   				jsii.String("accountIds"),
//   			},
//   			ResourceTags: []interface{}{
//   				[]interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analysisrule.html
//
type CfnAnalyzerPropsMixin_AnalysisRuleProperty struct {
	// A list of rules for the analyzer containing criteria to exclude from analysis.
	//
	// Entities that meet the rule criteria will not generate findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analysisrule.html#cfn-accessanalyzer-analyzer-analysisrule-exclusions
	//
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
}

