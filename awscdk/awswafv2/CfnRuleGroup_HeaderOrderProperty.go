package awswafv2


// The string containing the list of a web request's header names, ordered as they appear in the web request, separated by colons.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerOrderProperty := &HeaderOrderProperty{
//   	OversizeHandling: jsii.String("oversizeHandling"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-headerorder.html
//
type CfnRuleGroup_HeaderOrderProperty struct {
	// Handling of requests containing oversize fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-headerorder.html#cfn-wafv2-rulegroup-headerorder-oversizehandling
	//
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

