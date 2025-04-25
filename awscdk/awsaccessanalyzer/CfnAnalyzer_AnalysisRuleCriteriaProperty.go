package awsaccessanalyzer


// The criteria for an analysis rule for an analyzer.
//
// The criteria determine which entities will generate findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisRuleCriteriaProperty := &AnalysisRuleCriteriaProperty{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	ResourceTags: []interface{}{
//   		[]interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analysisrulecriteria.html
//
type CfnAnalyzer_AnalysisRuleCriteriaProperty struct {
	// A list of AWS account IDs to apply to the analysis rule criteria.
	//
	// The accounts cannot include the organization analyzer owner account. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers. The list cannot include more than 2,000 account IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analysisrulecriteria.html#cfn-accessanalyzer-analyzer-analysisrulecriteria-accountids
	//
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// An array of key-value pairs to match for your resources.
	//
	// You can use the set of Unicode letters, digits, whitespace, `_` , `.` , `/` , `=` , `+` , and `-` .
	//
	// For the tag key, you can specify a value that is 1 to 128 characters in length and cannot be prefixed with `aws:` .
	//
	// For the tag value, you can specify a value that is 0 to 256 characters in length. If the specified tag value is 0 characters, the rule is applied to all principals with the specified tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analysisrulecriteria.html#cfn-accessanalyzer-analyzer-analysisrulecriteria-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

