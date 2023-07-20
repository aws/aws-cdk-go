package awsaccessanalyzer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAnalyzer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalyzerProps := &CfnAnalyzerProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AnalyzerName: jsii.String("analyzerName"),
//   	ArchiveRules: []interface{}{
//   		&ArchiveRuleProperty{
//   			Filter: []interface{}{
//   				&FilterProperty{
//   					Property: jsii.String("property"),
//
//   					// the properties below are optional
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
//   				},
//   			},
//   			RuleName: jsii.String("ruleName"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html
//
type CfnAnalyzerProps struct {
	// The type represents the zone of trust for the analyzer.
	//
	// *Allowed Values* : ACCOUNT | ORGANIZATION.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the analyzer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-analyzername
	//
	AnalyzerName *string `field:"optional" json:"analyzerName" yaml:"analyzerName"`
	// Specifies the archive rules to add for the analyzer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-archiverules
	//
	ArchiveRules interface{} `field:"optional" json:"archiveRules" yaml:"archiveRules"`
	// The tags to apply to the analyzer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

