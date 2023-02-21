package awscloudfront


// Session stickiness provides the ability to define multiple requests from a single viewer as a single session.
//
// This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionStickinessConfigProperty := &sessionStickinessConfigProperty{
//   	idleTtl: jsii.Number(123),
//   	maximumTtl: jsii.Number(123),
//   }
//
type CfnContinuousDeploymentPolicy_SessionStickinessConfigProperty struct {
	// The amount of time after which you want sessions to cease if no requests are received.
	//
	// Allowed values are 300–3600 seconds (5–60 minutes).
	IdleTtl *float64 `field:"required" json:"idleTtl" yaml:"idleTtl"`
	// The maximum amount of time to consider requests from the viewer as being part of the same session.
	//
	// Allowed values are 300–3600 seconds (5–60 minutes).
	MaximumTtl *float64 `field:"required" json:"maximumTtl" yaml:"maximumTtl"`
}

