// An experiment to bundle the entire CDK into a single module
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
//   cfnRuleAssertion := &cfnRuleAssertion{
//   	assert: cfnConditionExpression,
//   	assertDescription: jsii.String("assertDescription"),
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

