package awselasticloadbalancingv2


// Information about the target group stickiness for a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupStickinessConfigProperty := &targetGroupStickinessConfigProperty{
//   	durationSeconds: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnListener_TargetGroupStickinessConfigProperty struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

