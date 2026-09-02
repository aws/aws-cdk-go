package interfacesawsdevopsagent


// A reference to a Trigger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerReference := &TriggerReference{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	TriggerArn: jsii.String("triggerArn"),
//   	TriggerId: jsii.String("triggerId"),
//   }
//
type TriggerReference struct {
	// The AgentSpaceId of the Trigger resource.
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The ARN of the Trigger resource.
	TriggerArn *string `field:"required" json:"triggerArn" yaml:"triggerArn"`
	// The TriggerId of the Trigger resource.
	TriggerId *string `field:"required" json:"triggerId" yaml:"triggerId"`
}

