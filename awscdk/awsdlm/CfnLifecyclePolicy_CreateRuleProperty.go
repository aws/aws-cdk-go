package awsdlm


// *[Custom snapshot and AMI policies only]* Specifies when the policy should create snapshots or AMIs.
//
// > - You must specify either *CronExpression* , or *Interval* , *IntervalUnit* , and *Times* .
// > - If you need to specify an [ArchiveRule](https://docs.aws.amazon.com/dlm/latest/APIReference/API_ArchiveRule.html) for the schedule, then you must specify a creation frequency of at least 28 days.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createRuleProperty := &CreateRuleProperty{
//   	CronExpression: jsii.String("cronExpression"),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   	Location: jsii.String("location"),
//   	Scripts: []interface{}{
//   		&ScriptProperty{
//   			ExecuteOperationOnScriptFailure: jsii.Boolean(false),
//   			ExecutionHandler: jsii.String("executionHandler"),
//   			ExecutionHandlerService: jsii.String("executionHandlerService"),
//   			ExecutionTimeout: jsii.Number(123),
//   			MaximumRetryCount: jsii.Number(123),
//   			Stages: []*string{
//   				jsii.String("stages"),
//   			},
//   		},
//   	},
//   	Times: []*string{
//   		jsii.String("times"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html
//
type CfnLifecyclePolicy_CreateRuleProperty struct {
	// The schedule, as a Cron expression.
	//
	// The schedule interval must be between 1 hour and 1 year. For more information, see the [Cron expressions reference](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cron-expressions.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-cronexpression
	//
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// The interval between snapshots.
	//
	// The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The interval unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-intervalunit
	//
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// *[Custom snapshot policies only]* Specifies the destination for snapshots created by the policy.
	//
	// The allowed destinations depend on the location of the targeted resources.
	//
	// - If the policy targets resources in a Region, then you must create snapshots in the same Region as the source resource.
	// - If the policy targets resources in a Local Zone, you can create snapshots in the same Local Zone or in its parent Region.
	// - If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost or in its parent Region.
	//
	// Specify one of the following values:
	//
	// - To create snapshots in the same Region as the source resource, specify `CLOUD` .
	// - To create snapshots in the same Local Zone as the source resource, specify `LOCAL_ZONE` .
	// - To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL` .
	//
	// Default: `CLOUD`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// *[Custom snapshot policies that target instances only]* Specifies pre and/or post scripts for a snapshot lifecycle policy that targets instances.
	//
	// This is useful for creating application-consistent snapshots, or for performing specific administrative tasks before or after Amazon Data Lifecycle Manager initiates snapshot creation.
	//
	// For more information, see [Automating application-consistent snapshots with pre and post scripts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/automate-app-consistent-backups.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-scripts
	//
	Scripts interface{} `field:"optional" json:"scripts" yaml:"scripts"`
	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon Data Lifecycle Manager selects a time within the next 24 hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-createrule.html#cfn-dlm-lifecyclepolicy-createrule-times
	//
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

