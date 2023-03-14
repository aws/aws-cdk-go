package awsaccessanalyzer


// The criteria for an archive rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRuleProperty := &archiveRuleProperty{
//   	filter: []interface{}{
//   		&filterProperty{
//   			property: jsii.String("property"),
//
//   			// the properties below are optional
//   			contains: []*string{
//   				jsii.String("contains"),
//   			},
//   			eq: []*string{
//   				jsii.String("eq"),
//   			},
//   			exists: jsii.Boolean(false),
//   			neq: []*string{
//   				jsii.String("neq"),
//   			},
//   		},
//   	},
//   	ruleName: jsii.String("ruleName"),
//   }
//
type CfnAnalyzer_ArchiveRuleProperty struct {
	// The criteria for the rule.
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// The name of the archive rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
}

