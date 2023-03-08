package awsroute53recoverycontrol


// A gating rule verifies that a gating routing control or set of gating routing controls, evaluates as true, based on a rule configuration that you specify, which allows a set of routing control state changes to complete.
//
// For example, if you specify one gating routing control and you set the `Type` in the rule configuration to `OR` , that indicates that you must set the gating routing control to `On` for the rule to evaluate as true; that is, for the gating control "switch" to be "On". When you do that, then you can update the routing control states for the target routing controls that you specify in the gating rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatingRuleProperty := &gatingRuleProperty{
//   	gatingControls: []*string{
//   		jsii.String("gatingControls"),
//   	},
//   	targetControls: []*string{
//   		jsii.String("targetControls"),
//   	},
//   	waitPeriodMs: jsii.Number(123),
//   }
//
type CfnSafetyRule_GatingRuleProperty struct {
	// An array of gating routing control Amazon Resource Names (ARNs).
	//
	// For a simple "on/off" switch, specify the ARN for one routing control. The gating routing controls are evaluated by the rule configuration that you specify to determine if the target routing control states can be changed.
	GatingControls *[]*string `field:"required" json:"gatingControls" yaml:"gatingControls"`
	// An array of target routing control Amazon Resource Names (ARNs) for which the states can only be updated if the rule configuration that you specify evaluates to true for the gating routing control.
	//
	// As a simple example, if you have a single gating control, it acts as an overall "on/off" switch for a set of target routing controls. You can use this to manually override automated failover, for example.
	TargetControls *[]*string `field:"required" json:"targetControls" yaml:"targetControls"`
	// An evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	//
	// This helps prevent "flapping" of state. The wait period is 5000 ms by default, but you can choose a custom value.
	WaitPeriodMs *float64 `field:"required" json:"waitPeriodMs" yaml:"waitPeriodMs"`
}

