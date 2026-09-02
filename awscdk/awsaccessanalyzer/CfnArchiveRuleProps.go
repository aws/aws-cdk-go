package awsaccessanalyzer


// Properties for defining a `CfnArchiveRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnArchiveRuleProps := &CfnArchiveRuleProps{
//   	AnalyzerName: jsii.String("analyzerName"),
//   	Filter: map[string]interface{}{
//   		"filterKey": &FilterItemsProperty{
//   			"contains": []*string{
//   				jsii.String("contains"),
//   			},
//   			"eq": []*string{
//   				jsii.String("eq"),
//   			},
//   			"exists": jsii.Boolean(false),
//   			"neq": []*string{
//   				jsii.String("neq"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-archiverule.html
//
type CfnArchiveRuleProps struct {
	// The name of the analyzer for the archive rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-archiverule.html#cfn-accessanalyzer-archiverule-analyzername
	//
	AnalyzerName *string `field:"required" json:"analyzerName" yaml:"analyzerName"`
	// The criteria for the archive rule.
	//
	// A map of filter criteria property names to their criterion values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-archiverule.html#cfn-accessanalyzer-archiverule-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// The name of the archive rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-archiverule.html#cfn-accessanalyzer-archiverule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
}

