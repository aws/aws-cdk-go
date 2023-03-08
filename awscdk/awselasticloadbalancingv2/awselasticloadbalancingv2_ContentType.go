package awselasticloadbalancingv2


// The content type for a fixed response.
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

