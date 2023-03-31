package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for `ListenerAction.forward()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   forwardOptions := &forwardOptions{
//   	stickinessDuration: duration,
//   }
//
// Experimental.
type ForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	// Experimental.
	StickinessDuration awscdk.Duration `field:"optional" json:"stickinessDuration" yaml:"stickinessDuration"`
}

