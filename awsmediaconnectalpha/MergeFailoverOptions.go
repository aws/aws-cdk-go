package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for merge-mode source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   mergeFailoverOptions := &MergeFailoverOptions{
//   	RecoveryWindow: cdk.Duration_Minutes(jsii.Number(30)),
//   	State: mediaconnect_alpha.State_ENABLED,
//   }
//
// Experimental.
type MergeFailoverOptions struct {
	// Search window time to look for SMPTE 2022-7 packets.
	//
	// A larger recovery window
	// means a longer delay in transmitting the stream, but more room for error
	// correction. A smaller window means a shorter delay but less room for correction.
	//
	// Valid range: 100–15000 ms.
	// Default: - 200 ms.
	//
	// Experimental.
	RecoveryWindow awscdk.Duration `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// Whether failover is enabled.
	//
	// Set to `State.DISABLED` to keep the configuration
	// on the flow without switching failover on.
	// Default: State.ENABLED
	//
	// Experimental.
	State State `field:"optional" json:"state" yaml:"state"`
}

