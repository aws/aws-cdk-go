package previewawsmediaconnectmixins


// The settings for source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   failoverConfigProperty := &FailoverConfigProperty{
//   	FailoverMode: jsii.String("failoverMode"),
//   	RecoveryWindow: jsii.Number(123),
//   	SourcePriority: &SourcePriorityProperty{
//   		PrimarySource: jsii.String("primarySource"),
//   	},
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-failoverconfig.html
//
type CfnFlowPropsMixin_FailoverConfigProperty struct {
	// The type of failover you choose for this flow.
	//
	// MERGE combines the source streams into a single stream, allowing graceful recovery from any single-source loss. FAILOVER allows switching between different streams. The string for this property must be entered as MERGE or FAILOVER. No other string entry is valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-failoverconfig.html#cfn-mediaconnect-flow-failoverconfig-failovermode
	//
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// Search window time to look for dash-7 packets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-failoverconfig.html#cfn-mediaconnect-flow-failoverconfig-recoverywindow
	//
	RecoveryWindow *float64 `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// The priority you want to assign to a source.
	//
	// You can have a primary stream and a backup stream or two equally prioritized streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-failoverconfig.html#cfn-mediaconnect-flow-failoverconfig-sourcepriority
	//
	SourcePriority interface{} `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// The state of source failover on the flow.
	//
	// If the state is inactive, the flow can have only one source. If the state is active, the flow can have one or two sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-failoverconfig.html#cfn-mediaconnect-flow-failoverconfig-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

