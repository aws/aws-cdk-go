package awswafv2


// The path component of the URI Fragment.
//
// This is the part of a web request that identifies a fragment uri, for example, /abcd#introduction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uriFragmentProperty := &UriFragmentProperty{
//   	FallbackBehavior: jsii.String("fallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-urifragment.html
//
type CfnRuleGroup_UriFragmentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-urifragment.html#cfn-wafv2-rulegroup-urifragment-fallbackbehavior
	//
	FallbackBehavior *string `field:"optional" json:"fallbackBehavior" yaml:"fallbackBehavior"`
}

