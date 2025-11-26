package previewawsdlmmixins


// *[Event-based policies only]* Specifies an event that activates an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eventSourceProperty := &EventSourceProperty{
//   	Parameters: &EventParametersProperty{
//   		DescriptionRegex: jsii.String("descriptionRegex"),
//   		EventType: jsii.String("eventType"),
//   		SnapshotOwner: []*string{
//   			jsii.String("snapshotOwner"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventsource.html
//
type CfnLifecyclePolicyPropsMixin_EventSourceProperty struct {
	// Information about the event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventsource.html#cfn-dlm-lifecyclepolicy-eventsource-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The source of the event.
	//
	// Currently only managed Amazon EventBridge (formerly known as Amazon CloudWatch) events are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventsource.html#cfn-dlm-lifecyclepolicy-eventsource-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

