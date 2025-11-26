package previewawsgameliftmixins


// Rule that controls how a fleet is scaled.
//
// Scaling policies are uniquely identified by the combination of name and fleet ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingPolicyProperty := &ScalingPolicyProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	EvaluationPeriods: jsii.Number(123),
//   	Location: jsii.String("location"),
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	PolicyType: jsii.String("policyType"),
//   	ScalingAdjustment: jsii.Number(123),
//   	ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   	Status: jsii.String("status"),
//   	TargetConfiguration: &TargetConfigurationProperty{
//   		TargetValue: jsii.Number(123),
//   	},
//   	Threshold: jsii.Number(123),
//   	UpdateStatus: jsii.String("updateStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html
//
type CfnFleetPropsMixin_ScalingPolicyProperty struct {
	// Comparison operator to use when measuring a metric against the threshold value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-evaluationperiods
	//
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The fleet location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Name of the Amazon GameLift Servers-defined metric that is used to trigger a scaling adjustment.
	//
	// For detailed descriptions of fleet metrics, see [Monitor Amazon GameLift Servers with Amazon CloudWatch](https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html) .
	//
	// - *ActivatingGameSessions* -- Game sessions in the process of being created.
	// - *ActiveGameSessions* -- Game sessions that are currently running.
	// - *ActiveInstances* -- Fleet instances that are currently running at least one game session.
	// - *AvailableGameSessions* -- Additional game sessions that fleet could host simultaneously, given current capacity.
	// - *AvailablePlayerSessions* -- Empty player slots in currently active game sessions. This includes game sessions that are not currently accepting players. Reserved player slots are not included.
	// - *CurrentPlayerSessions* -- Player slots in active game sessions that are being used by a player or are reserved for a player.
	// - *IdleInstances* -- Active instances that are currently hosting zero game sessions.
	// - *PercentAvailableGameSessions* -- Unused percentage of the total number of game sessions that a fleet could host simultaneously, given current capacity. Use this metric for a target-based scaling policy.
	// - *PercentIdleInstances* -- Percentage of the total number of active instances that are hosting zero game sessions.
	// - *QueueDepth* -- Pending game session placement requests, in any queue, where the current fleet is the top-priority destination.
	// - *WaitTime* -- Current wait time for pending game session placement requests, in any queue, where the current fleet is the top-priority destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A descriptive label that is associated with a fleet's scaling policy.
	//
	// Policy names do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of scaling policy to create.
	//
	// For a target-based policy, set the parameter *MetricName* to 'PercentAvailableGameSessions' and specify a *TargetConfiguration* . For a rule-based policy set the following parameters: *MetricName* , *ComparisonOperator* , *Threshold* , *EvaluationPeriods* , *ScalingAdjustmentType* , and *ScalingAdjustment* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Amount of adjustment to make, based on the scaling adjustment type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-scalingadjustment
	//
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The type of adjustment to make to a fleet's instance count.
	//
	// - *ChangeInCapacity* -- add (or subtract) the scaling adjustment value from the current instance count. Positive values scale up while negative values scale down.
	// - *ExactCapacity* -- set the instance count to the scaling adjustment value.
	// - *PercentChangeInCapacity* -- increase or reduce the current instance count by the scaling adjustment, read as a percentage. Positive values scale up while negative values scale down.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-scalingadjustmenttype
	//
	ScalingAdjustmentType *string `field:"optional" json:"scalingAdjustmentType" yaml:"scalingAdjustmentType"`
	// Current status of the scaling policy.
	//
	// The scaling policy can be in force only when in an `ACTIVE` status. Scaling policies can be suspended for individual fleets. If the policy is suspended for a fleet, the policy status does not change.
	//
	// - *ACTIVE* -- The scaling policy can be used for auto-scaling a fleet.
	// - *UPDATE_REQUESTED* -- A request to update the scaling policy has been received.
	// - *UPDATING* -- A change is being made to the scaling policy.
	// - *DELETE_REQUESTED* -- A request to delete the scaling policy has been received.
	// - *DELETING* -- The scaling policy is being deleted.
	// - *DELETED* -- The scaling policy has been deleted.
	// - *ERROR* -- An error occurred in creating the policy. It should be removed and recreated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An object that contains settings for a target-based scaling policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-targetconfiguration
	//
	TargetConfiguration interface{} `field:"optional" json:"targetConfiguration" yaml:"targetConfiguration"`
	// Metric value used to trigger a scaling event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// The current status of the fleet's scaling policies in a requested fleet location.
	//
	// The status `PENDING_UPDATE` indicates that an update was requested for the fleet but has not yet been completed for the location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-scalingpolicy.html#cfn-gamelift-fleet-scalingpolicy-updatestatus
	//
	UpdateStatus *string `field:"optional" json:"updateStatus" yaml:"updateStatus"`
}

