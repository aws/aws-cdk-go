package awsaccessanalyzer


// The criteria that defines the archive rule.
//
// To learn about filter keys that you can use to create an archive rule, see [filter keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html) in the *User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Property: jsii.String("property"),
//
//   	// the properties below are optional
//   	Contains: []*string{
//   		jsii.String("contains"),
//   	},
//   	Eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	Exists: jsii.Boolean(false),
//   	Neq: []*string{
//   		jsii.String("neq"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html
//
type CfnAnalyzer_FilterProperty struct {
	// The property used to define the criteria in the filter for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-property
	//
	Property *string `field:"required" json:"property" yaml:"property"`
	// A "contains" condition to match for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-contains
	//
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// An "equals" condition to match for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-eq
	//
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// An "exists" condition to match for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-exists
	//
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// A "not equal" condition to match for the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-filter.html#cfn-accessanalyzer-analyzer-filter-neq
	//
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

