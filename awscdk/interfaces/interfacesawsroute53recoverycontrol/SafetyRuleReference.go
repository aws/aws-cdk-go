package interfacesawsroute53recoverycontrol


// A reference to a SafetyRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   safetyRuleReference := &SafetyRuleReference{
//   	SafetyRuleArn: jsii.String("safetyRuleArn"),
//   }
//
type SafetyRuleReference struct {
	// The SafetyRuleArn of the SafetyRule resource.
	SafetyRuleArn *string `field:"required" json:"safetyRuleArn" yaml:"safetyRuleArn"`
}

