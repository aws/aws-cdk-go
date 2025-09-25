package awsobservabilityadmin


// A reference to a TelemetryRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryRuleReference := &TelemetryRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type TelemetryRuleReference struct {
	// The RuleArn of the TelemetryRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

