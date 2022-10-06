package awselasticloadbalancingv2


// Options for `ListenerAction.fixedResponse()`.
//
// Example:
//   var listener applicationListener
//
//
//   listener.addAction(jsii.String("Fixed"), &addApplicationActionProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	action: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   		contentType: elbv2.contentType_TEXT_PLAIN,
//   		messageBody: jsii.String("OK"),
//   	}),
//   })
//
type FixedResponseOptions struct {
	// Content Type of the response.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The response body.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

