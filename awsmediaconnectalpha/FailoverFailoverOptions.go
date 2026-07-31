package awsmediaconnectalpha


// Options for switchover-mode source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   failoverFailoverOptions := &FailoverFailoverOptions{
//   	PrimarySource: jsii.String("primarySource"),
//   	State: mediaconnect_alpha.State_ENABLED,
//   }
//
// Experimental.
type FailoverFailoverOptions struct {
	// The name of the source you want to treat as primary.
	//
	// If set, MediaConnect always
	// uses this source when it is available. When unset, both sources are treated with
	// equal priority.
	// Default: - both sources are equal priority.
	//
	// Experimental.
	PrimarySource *string `field:"optional" json:"primarySource" yaml:"primarySource"`
	// Whether failover is enabled.
	//
	// Set to `State.DISABLED` to keep the configuration
	// on the flow without switching failover on.
	// Default: State.ENABLED
	//
	// Experimental.
	State State `field:"optional" json:"state" yaml:"state"`
}

