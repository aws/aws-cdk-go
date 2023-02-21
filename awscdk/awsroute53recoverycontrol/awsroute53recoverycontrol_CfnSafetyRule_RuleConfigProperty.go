package awsroute53recoverycontrol


// The rule configuration for an assertion rule.
//
// That is, the criteria that you set for specific assertion controls (routing controls) that specify how many control states must be `ON` after a transaction completes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConfigProperty := &ruleConfigProperty{
//   	inverted: jsii.Boolean(false),
//   	threshold: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnSafetyRule_RuleConfigProperty struct {
	// Logical negation of the rule.
	//
	// If the rule would usually evaluate true, it's evaluated as false, and vice versa.
	Inverted interface{} `field:"required" json:"inverted" yaml:"inverted"`
	// The value of N, when you specify an `ATLEAST` rule type.
	//
	// That is, `Threshold` is the number of controls that must be set when you specify an `ATLEAST` type.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// A rule can be one of the following: `ATLEAST` , `AND` , or `OR` .
	Type *string `field:"required" json:"type" yaml:"type"`
}

