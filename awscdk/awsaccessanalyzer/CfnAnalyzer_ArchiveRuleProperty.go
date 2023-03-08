package awsaccessanalyzer


// The criteria for an archive rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRuleProperty := &ArchiveRuleProperty{
//   	Filter: []interface{}{
//   		&FilterProperty{
//   			Property: jsii.String("property"),
//
//   			// the properties below are optional
//   			Contains: []*string{
//   				jsii.String("contains"),
//   			},
//   			Eq: []*string{
//   				jsii.String("eq"),
//   			},
//   			Exists: jsii.Boolean(false),
//   			Neq: []*string{
//   				jsii.String("neq"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }
//
type CfnAnalyzer_ArchiveRuleProperty struct {
	// The criteria for the rule.
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// The name of the archive rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
}

