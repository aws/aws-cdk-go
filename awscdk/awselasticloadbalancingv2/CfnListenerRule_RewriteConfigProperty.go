package awselasticloadbalancingv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rewriteConfigProperty := &RewriteConfigProperty{
//   	Regex: jsii.String("regex"),
//   	Replace: jsii.String("replace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html
//
type CfnListenerRule_RewriteConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfig-regex
	//
	Regex *string `field:"required" json:"regex" yaml:"regex"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfig.html#cfn-elasticloadbalancingv2-listenerrule-rewriteconfig-replace
	//
	Replace *string `field:"required" json:"replace" yaml:"replace"`
}

