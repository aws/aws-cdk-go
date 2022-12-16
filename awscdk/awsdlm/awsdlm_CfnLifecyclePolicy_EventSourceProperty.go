package awsdlm


// Specifies an event that triggers an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &eventSourceProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	parameters: &eventParametersProperty{
//   		eventType: jsii.String("eventType"),
//   		snapshotOwner: []*string{
//   			jsii.String("snapshotOwner"),
//   		},
//
//   		// the properties below are optional
//   		descriptionRegex: jsii.String("descriptionRegex"),
//   	},
//   }
//
type CfnLifecyclePolicy_EventSourceProperty struct {
	// The source of the event.
	//
	// Currently only managed CloudWatch Events rules are supported.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Information about the event.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

