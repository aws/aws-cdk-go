package awsaccessanalyzer


// The criteria that defines the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	property: jsii.String("property"),
//
//   	// the properties below are optional
//   	contains: []*string{
//   		jsii.String("contains"),
//   	},
//   	eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	exists: jsii.Boolean(false),
//   	neq: []*string{
//   		jsii.String("neq"),
//   	},
//   }
//
type CfnAnalyzer_FilterProperty struct {
	// The property used to define the criteria in the filter for the rule.
	Property *string `field:"required" json:"property" yaml:"property"`
	// A "contains" condition to match for the rule.
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// An "equals" condition to match for the rule.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// An "exists" condition to match for the rule.
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// A "not equal" condition to match for the rule.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

