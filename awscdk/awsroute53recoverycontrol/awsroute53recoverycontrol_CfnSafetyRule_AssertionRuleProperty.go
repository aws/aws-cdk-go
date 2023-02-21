package awsroute53recoverycontrol


// An assertion rule enforces that, when you change a routing control state, that the criteria that you set in the rule configuration is met.
//
// Otherwise, the change to the routing control is not accepted. For example, the criteria might be that at least one routing control state is `On` after the transaction so that traffic continues to flow to at least one cell for the application. This ensures that you avoid a fail-open scenario.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionRuleProperty := &assertionRuleProperty{
//   	assertedControls: []*string{
//   		jsii.String("assertedControls"),
//   	},
//   	waitPeriodMs: jsii.Number(123),
//   }
//
type CfnSafetyRule_AssertionRuleProperty struct {
	// The routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed.
	//
	// For example, you might include three routing controls, one for each of three AWS Regions.
	AssertedControls *[]*string `field:"required" json:"assertedControls" yaml:"assertedControls"`
	// An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	//
	// This helps prevent "flapping" of state. The wait period is 5000 ms by default, but you can choose a custom value.
	WaitPeriodMs *float64 `field:"required" json:"waitPeriodMs" yaml:"waitPeriodMs"`
}

