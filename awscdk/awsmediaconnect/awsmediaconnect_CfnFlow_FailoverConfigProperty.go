package awsmediaconnect


// The settings for source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConfigProperty := &failoverConfigProperty{
//   	failoverMode: jsii.String("failoverMode"),
//   	recoveryWindow: jsii.Number(123),
//   	sourcePriority: &sourcePriorityProperty{
//   		primarySource: jsii.String("primarySource"),
//   	},
//   	state: jsii.String("state"),
//   }
//
type CfnFlow_FailoverConfigProperty struct {
	// The type of failover you choose for this flow.
	//
	// MERGE combines the source streams into a single stream, allowing graceful recovery from any single-source loss. FAILOVER allows switching between different streams.
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// The size of the buffer (delay) that the service maintains.
	//
	// A larger buffer means a longer delay in transmitting the stream, but more room for error correction. A smaller buffer means a shorter delay, but less room for error correction. You can choose a value from 100-500 ms. If you keep this field blank, the service uses the default value of 200 ms. This setting only applies when Failover Mode is set to MERGE.
	RecoveryWindow *float64 `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// The priority you want to assign to a source.
	//
	// You can have a primary stream and a backup stream or two equally prioritized streams. This setting only applies when Failover Mode is set to FAILOVER.
	SourcePriority interface{} `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// The state of source failover on the flow.
	//
	// If the state is disabled, the flow can have only one source. If the state is enabled, the flow can have one or two sources.
	State *string `field:"optional" json:"state" yaml:"state"`
}

