package mixinsawselasticloadbalancingv2


// Information about a rewrite transform.
//
// This transform matches a pattern and replaces it with the specified string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rewriteConfigProperty := &RewriteConfigProperty{
//   	Regex: jsii.String("regex"),
//   	Replace: jsii.String("replace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html
//
type CfnListenerRulePropsMixin_RewriteConfigProperty struct {
	// The regular expression to match in the input string.
	//
	// The maximum length of the string is 1,024 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfig-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// The replacement string to use when rewriting the matched input.
	//
	// The maximum length of the string is 1,024 characters. You can specify capture groups in the regular expression (for example, $1 and $2).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfig-replace
	//
	Replace *string `field:"optional" json:"replace" yaml:"replace"`
}

