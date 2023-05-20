package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The queue setting that determines the highest latency allowed for individual players when placing a game session.
//
// When a latency policy is in force, a game session cannot be placed with any fleet in a Region where a player reports latency higher than the cap.
//
// Latency policies are only enforced when the placement request contains player latency information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   playerLatencyPolicy := &PlayerLatencyPolicy{
//   	MaximumIndividualPlayerLatency: cdk.Duration_Minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	PolicyDuration: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type PlayerLatencyPolicy struct {
	// The maximum latency value that is allowed for any player, in milliseconds.
	//
	// All policies must have a value set for this property.
	// Experimental.
	MaximumIndividualPlayerLatency awscdk.Duration `field:"required" json:"maximumIndividualPlayerLatency" yaml:"maximumIndividualPlayerLatency"`
	// The length of time, in seconds, that the policy is enforced while placing a new game session.
	// Experimental.
	PolicyDuration awscdk.Duration `field:"optional" json:"policyDuration" yaml:"policyDuration"`
}

