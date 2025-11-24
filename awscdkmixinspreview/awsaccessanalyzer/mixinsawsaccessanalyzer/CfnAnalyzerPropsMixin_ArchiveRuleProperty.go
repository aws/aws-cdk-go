package mixinsawsaccessanalyzer


// Contains information about an archive rule.
//
// Archive rules automatically archive new findings that meet the criteria you define when you create the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveRuleProperty := &ArchiveRuleProperty{
//   	Filter: []interface{}{
//   		&FilterProperty{
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
//   			Property: jsii.String("property"),
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html
//
type CfnAnalyzerPropsMixin_ArchiveRuleProperty struct {
	// The criteria for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// The name of the rule to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-archiverule.html#cfn-accessanalyzer-analyzer-archiverule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

