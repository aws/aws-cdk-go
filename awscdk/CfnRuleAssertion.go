package awscdk


// A rule assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//
//   cfnRuleAssertion := &CfnRuleAssertion{
//   	Assert: cfnConditionExpression,
//   	AssertDescription: jsii.String("assertDescription"),
//   }
//
// Experimental.
type CfnRuleAssertion struct {
	// The assertion.
	// Experimental.
	Assert ICfnConditionExpression `field:"required" json:"assert" yaml:"assert"`
	// The assertion description.
	// Experimental.
	AssertDescription *string `field:"required" json:"assertDescription" yaml:"assertDescription"`
}

