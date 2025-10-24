package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// An assertion within a template rule, defined by intrinsic functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnRuleConditionExpression ICfnRuleConditionExpression
//
//   templateRuleAssertion := &TemplateRuleAssertion{
//   	Assert: cfnRuleConditionExpression,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type TemplateRuleAssertion struct {
	// The assertion condition.
	Assert awscdk.ICfnRuleConditionExpression `field:"required" json:"assert" yaml:"assert"`
	// The description for the asssertion.
	// Default: - no description provided for the assertion.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

