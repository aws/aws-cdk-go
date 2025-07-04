package awswafv2


// A single regular expression.
//
// This is used in a `RegexPatternSet` and also in the configuration for the AWS Managed Rules rule group `AWSManagedRulesAntiDDoSRuleSet` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regexProperty := &RegexProperty{
//   	RegexString: jsii.String("regexString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regex.html
//
type CfnWebACL_RegexProperty struct {
	// The string representing the regular expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-regex.html#cfn-wafv2-webacl-regex-regexstring
	//
	RegexString *string `field:"optional" json:"regexString" yaml:"regexString"`
}

