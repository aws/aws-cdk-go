package awspipes


// The parameters for using an EventBridge event bus as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetEventBridgeEventBusParametersProperty := &pipeTargetEventBridgeEventBusParametersProperty{
//   	detailType: jsii.String("detailType"),
//   	endpointId: jsii.String("endpointId"),
//   	resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	source: jsii.String("source"),
//   	time: jsii.String("time"),
//   }
//
type CfnPipe_PipeTargetEventBridgeEventBusParametersProperty struct {
	// A free-form string, with a maximum of 128 characters, used to decide what fields to expect in the event detail.
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// The URL subdomain of the endpoint.
	//
	// For example, if the URL for Endpoint is https://abcde.veo.endpoints.event.amazonaws.com, then the EndpointId is `abcde.veo` .
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// AWS resources, identified by Amazon Resource Name (ARN), which the event primarily concerns.
	//
	// Any number, including zero, may be present.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The source of the event.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The time stamp of the event, per [RFC3339](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc3339.txt) . If no time stamp is provided, the time stamp of the [PutEvents](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) call is used.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

