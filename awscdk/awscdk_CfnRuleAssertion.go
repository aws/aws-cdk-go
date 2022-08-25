// Version 2 of the AWS Cloud Development Kit library
package awscdk


// A rule assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnConditionExpression iCfnConditionExpression
//
//   cfnRuleAssertion := &cfnRuleAssertion{
//   	assert: cfnConditionExpression,
//   	assertDescription: jsii.String("assertDescription"),
//   }
//
type CfnRuleAssertion struct {
	// The assertion.
	Assert ICfnConditionExpression `field:"required" json:"assert" yaml:"assert"`
	// The assertion description.
	AssertDescription *string `field:"required" json:"assertDescription" yaml:"assertDescription"`
}

