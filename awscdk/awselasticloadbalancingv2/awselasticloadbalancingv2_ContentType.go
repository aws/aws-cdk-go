package awselasticloadbalancingv2


// The content type for a fixed response.
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
//   		ContentType: elbv2.ContentType_TEXT_PLAIN,
//   		MessageBody: jsii.String("OK"),
//   	}),
//   })
//
// Deprecated: superceded by `FixedResponseOptions`.
type ContentType string

const (
	// Deprecated: superceded by `FixedResponseOptions`.
	ContentType_TEXT_PLAIN ContentType = "TEXT_PLAIN"
	// Deprecated: superceded by `FixedResponseOptions`.
	ContentType_TEXT_CSS ContentType = "TEXT_CSS"
	// Deprecated: superceded by `FixedResponseOptions`.
	ContentType_TEXT_HTML ContentType = "TEXT_HTML"
	// Deprecated: superceded by `FixedResponseOptions`.
	ContentType_APPLICATION_JAVASCRIPT ContentType = "APPLICATION_JAVASCRIPT"
	// Deprecated: superceded by `FixedResponseOptions`.
	ContentType_APPLICATION_JSON ContentType = "APPLICATION_JSON"
)

