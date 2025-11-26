package previewawsaccessanalyzermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAnalyzerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnalyzerMixinProps := &CfnAnalyzerMixinProps{
//   	AnalyzerConfiguration: &AnalyzerConfigurationProperty{
//   		InternalAccessConfiguration: &InternalAccessConfigurationProperty{
//   			InternalAccessAnalysisRule: &InternalAccessAnalysisRuleProperty{
//   				Inclusions: []interface{}{
//   					&InternalAccessAnalysisRuleCriteriaProperty{
//   						AccountIds: []*string{
//   							jsii.String("accountIds"),
//   						},
//   						ResourceArns: []*string{
//   							jsii.String("resourceArns"),
//   						},
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		UnusedAccessConfiguration: &UnusedAccessConfigurationProperty{
//   			AnalysisRule: &AnalysisRuleProperty{
//   				Exclusions: []interface{}{
//   					&AnalysisRuleCriteriaProperty{
//   						AccountIds: []*string{
//   							jsii.String("accountIds"),
//   						},
//   						ResourceTags: []interface{}{
//   							[]interface{}{
//   								&CfnTag{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   			UnusedAccessAge: jsii.Number(123),
//   		},
//   	},
//   	AnalyzerName: jsii.String("analyzerName"),
//   	ArchiveRules: []interface{}{
//   		&ArchiveRuleProperty{
//   			Filter: []interface{}{
//   				&FilterProperty{
//   					Contains: []*string{
//   						jsii.String("contains"),
//   					},
//   					Eq: []*string{
//   						jsii.String("eq"),
//   					},
//   					Exists: jsii.Boolean(false),
//   					Neq: []*string{
//   						jsii.String("neq"),
//   					},
//   					Property: jsii.String("property"),
//   				},
//   			},
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html
//
type CfnAnalyzerMixinProps struct {
	// Contains information about the configuration of an analyzer for an AWS organization or account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-analyzerconfiguration
	//
	AnalyzerConfiguration interface{} `field:"optional" json:"analyzerConfiguration" yaml:"analyzerConfiguration"`
	// The name of the analyzer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-analyzername
	//
	AnalyzerName *string `field:"optional" json:"analyzerName" yaml:"analyzerName"`
	// Specifies the archive rules to add for the analyzer.
	//
	// Archive rules automatically archive findings that meet the criteria you define for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-archiverules
	//
	ArchiveRules interface{} `field:"optional" json:"archiveRules" yaml:"archiveRules"`
	// An array of key-value pairs to apply to the analyzer.
	//
	// You can use the set of Unicode letters, digits, whitespace, `_` , `.` , `/` , `=` , `+` , and `-` .
	//
	// For the tag key, you can specify a value that is 1 to 128 characters in length and cannot be prefixed with `aws:` .
	//
	// For the tag value, you can specify a value that is 0 to 256 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type represents the zone of trust for the analyzer.
	//
	// *Allowed Values* : ACCOUNT | ORGANIZATION | ACCOUNT_UNUSED_ACCESS | ACCOUNT_INTERNAL_ACCESS | ORGANIZATION_INTERNAL_ACCESS | ORGANIZATION_UNUSED_ACCESS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

