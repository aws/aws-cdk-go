package interfacesawsnotifications


// A reference to a EventRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventRuleReference := &EventRuleReference{
//   	EventRuleArn: jsii.String("eventRuleArn"),
//   }
//
type EventRuleReference struct {
	// The Arn of the EventRule resource.
	EventRuleArn *string `field:"required" json:"eventRuleArn" yaml:"eventRuleArn"`
}

