package awselasticloadbalancingv2


// Options for `ListenerAction.fixedResponse()`.
//
// Example:
//   var listener applicationListener
//
//
//   listener.AddAction(jsii.String("Fixed"), &AddApplicationActionProps{
//   	Priority: jsii.Number(10),
//   	Conditions: []listenerCondition{
//   		elbv2.*listenerCondition_PathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	Action: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
//   		ContentType: jsii.String("text/plain"),
//   		MessageBody: jsii.String("OK"),
//   	}),
//   })
//
type FixedResponseOptions struct {
	// Content Type of the response.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	// Default: - Automatically determined.
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The response body.
	// Default: - No body.
	//
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

