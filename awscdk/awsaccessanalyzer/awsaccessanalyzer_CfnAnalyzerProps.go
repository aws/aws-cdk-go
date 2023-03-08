package awsaccessanalyzer

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAnalyzer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalyzerProps := &cfnAnalyzerProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	analyzerName: jsii.String("analyzerName"),
//   	archiveRules: []interface{}{
//   		&archiveRuleProperty{
//   			filter: []interface{}{
//   				&filterProperty{
//   					property: jsii.String("property"),
//
//   					// the properties below are optional
//   					contains: []*string{
//   						jsii.String("contains"),
//   					},
//   					eq: []*string{
//   						jsii.String("eq"),
//   					},
//   					exists: jsii.Boolean(false),
//   					neq: []*string{
//   						jsii.String("neq"),
//   					},
//   				},
//   			},
//   			ruleName: jsii.String("ruleName"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAnalyzerProps struct {
	// The type represents the zone of trust for the analyzer.
	//
	// *Allowed Values* : ACCOUNT | ORGANIZATION.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of the analyzer.
	AnalyzerName *string `field:"optional" json:"analyzerName" yaml:"analyzerName"`
	// Specifies the archive rules to add for the analyzer.
	ArchiveRules interface{} `field:"optional" json:"archiveRules" yaml:"archiveRules"`
	// The tags to apply to the analyzer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

