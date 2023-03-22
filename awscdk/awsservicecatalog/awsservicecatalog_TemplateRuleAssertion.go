package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// An assertion within a template rule, defined by intrinsic functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnRuleConditionExpression iCfnRuleConditionExpression
//
//   templateRuleAssertion := &templateRuleAssertion{
//   	assert: cfnRuleConditionExpression,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
// Experimental.
type TemplateRuleAssertion struct {
	// The assertion condition.
	// Experimental.
	Assert awscdk.ICfnRuleConditionExpression `field:"required" json:"assert" yaml:"assert"`
	// The description for the asssertion.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

