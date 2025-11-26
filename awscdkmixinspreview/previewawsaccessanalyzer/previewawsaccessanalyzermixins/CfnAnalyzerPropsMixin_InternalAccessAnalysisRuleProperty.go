package previewawsaccessanalyzermixins


// Contains information about analysis rules for the internal access analyzer.
//
// Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   internalAccessAnalysisRuleProperty := &InternalAccessAnalysisRuleProperty{
//   	Inclusions: []interface{}{
//   		&InternalAccessAnalysisRuleCriteriaProperty{
//   			AccountIds: []*string{
//   				jsii.String("accountIds"),
//   			},
//   			ResourceArns: []*string{
//   				jsii.String("resourceArns"),
//   			},
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrule.html
//
type CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleProperty struct {
	// A list of rules for the internal access analyzer containing criteria to include in analysis.
	//
	// Only resources that meet the rule criteria will generate findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrule.html#cfn-accessanalyzer-analyzer-internalaccessanalysisrule-inclusions
	//
	Inclusions interface{} `field:"optional" json:"inclusions" yaml:"inclusions"`
}

