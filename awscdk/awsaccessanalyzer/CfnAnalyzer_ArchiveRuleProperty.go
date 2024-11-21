package awsaccessanalyzer


// Contains information about an archive rule.
//
// Archive rules automatically archive new findings that meet the criteria you define when you create the rule.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html
//
type CfnAnalyzer_ArchiveRuleProperty struct {
	// The criteria for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// The name of the rule to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
}

