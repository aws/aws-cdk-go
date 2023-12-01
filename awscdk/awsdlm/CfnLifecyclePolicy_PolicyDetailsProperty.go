package awsdlm


// Specifies the configuration of a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDetailsProperty := &PolicyDetailsProperty{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			CrossRegionCopy: []interface{}{
//   				&CrossRegionCopyActionProperty{
//   					EncryptionConfiguration: &EncryptionConfigurationProperty{
//   						Encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						CmkArn: jsii.String("cmkArn"),
//   					},
//   					Target: jsii.String("target"),
//
//   					// the properties below are optional
//   					RetainRule: &CrossRegionCopyRetainRuleProperty{
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	EventSource: &EventSourceProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Parameters: &EventParametersProperty{
//   			EventType: jsii.String("eventType"),
//   			SnapshotOwner: []*string{
//   				jsii.String("snapshotOwner"),
//   			},
//
//   			// the properties below are optional
//   			DescriptionRegex: jsii.String("descriptionRegex"),
//   		},
//   	},
//   	Parameters: &ParametersProperty{
//   		ExcludeBootVolume: jsii.Boolean(false),
//   		ExcludeDataVolumeTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NoReboot: jsii.Boolean(false),
//   	},
//   	PolicyType: jsii.String("policyType"),
//   	ResourceLocations: []*string{
//   		jsii.String("resourceLocations"),
//   	},
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	Schedules: []interface{}{
//   		&ScheduleProperty{
//   			ArchiveRule: &ArchiveRuleProperty{
//   				RetainRule: &ArchiveRetainRuleProperty{
//   					RetentionArchiveTier: &RetentionArchiveTierProperty{
//   						Count: jsii.Number(123),
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   				},
//   			},
//   			CopyTags: jsii.Boolean(false),
//   			CreateRule: &CreateRuleProperty{
//   				CronExpression: jsii.String("cronExpression"),
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   				Location: jsii.String("location"),
//   				Scripts: []interface{}{
//   					&ScriptProperty{
//   						ExecuteOperationOnScriptFailure: jsii.Boolean(false),
//   						ExecutionHandler: jsii.String("executionHandler"),
//   						ExecutionHandlerService: jsii.String("executionHandlerService"),
//   						ExecutionTimeout: jsii.Number(123),
//   						MaximumRetryCount: jsii.Number(123),
//   						Stages: []*string{
//   							jsii.String("stages"),
//   						},
//   					},
//   				},
//   				Times: []*string{
//   					jsii.String("times"),
//   				},
//   			},
//   			CrossRegionCopyRules: []interface{}{
//   				&CrossRegionCopyRuleProperty{
//   					Encrypted: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					CmkArn: jsii.String("cmkArn"),
//   					CopyTags: jsii.Boolean(false),
//   					DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   					RetainRule: &CrossRegionCopyRetainRuleProperty{
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   					Target: jsii.String("target"),
//   					TargetRegion: jsii.String("targetRegion"),
//   				},
//   			},
//   			DeprecateRule: &DeprecateRuleProperty{
//   				Count: jsii.Number(123),
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			FastRestoreRule: &FastRestoreRuleProperty{
//   				AvailabilityZones: []*string{
//   					jsii.String("availabilityZones"),
//   				},
//   				Count: jsii.Number(123),
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			Name: jsii.String("name"),
//   			RetainRule: &RetainRuleProperty{
//   				Count: jsii.Number(123),
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			ShareRules: []interface{}{
//   				&ShareRuleProperty{
//   					TargetAccounts: []*string{
//   						jsii.String("targetAccounts"),
//   					},
//   					UnshareInterval: jsii.Number(123),
//   					UnshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   				},
//   			},
//   			TagsToAdd: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			VariableTags: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TargetTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html
//
type CfnLifecyclePolicy_PolicyDetailsProperty struct {
	// *[Event-based policies only]* The actions to be performed when the event-based policy is activated.
	//
	// You can specify only one action per policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// *[Event-based policies only]* The event that activates the event-based policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-eventsource
	//
	EventSource interface{} `field:"optional" json:"eventSource" yaml:"eventSource"`
	// *[Custom snapshot and AMI policies only]* A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// > If you are modifying a policy that was created or previously modified using the Amazon Data Lifecycle Manager console, then you must include this parameter and specify either the default values or the new values that you require. You can't omit this parameter or set its values to null.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// *[Custom policies only]* The valid target resource types and actions a policy can manage.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account .
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// *[Custom snapshot and AMI policies only]* The location of the resources to backup.
	//
	// If the source resources are located in an AWS Region , specify `CLOUD` . If the source resources are located on an Outpost in your account, specify `OUTPOST` .
	//
	// If you specify `OUTPOST` , Amazon Data Lifecycle Manager backs up all resources of the specified type with matching target tags across all of the Outposts in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcelocations
	//
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// *[Custom snapshot policies only]* The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// *[Custom snapshot and AMI policies only]* The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	//
	// A policy can have up to four schedules—one mandatory schedule and up to three optional schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-schedules
	//
	Schedules interface{} `field:"optional" json:"schedules" yaml:"schedules"`
	// *[Custom snapshot and AMI policies only]* The single tag that identifies targeted resources for this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-targettags
	//
	TargetTags interface{} `field:"optional" json:"targetTags" yaml:"targetTags"`
}

