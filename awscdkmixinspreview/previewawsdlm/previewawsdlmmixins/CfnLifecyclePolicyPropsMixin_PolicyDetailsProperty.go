package previewawsdlmmixins


// Specifies the configuration of a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var crossRegionCopyTargets interface{}
//   var excludeTags interface{}
//   var excludeVolumeTypes interface{}
//
//   policyDetailsProperty := &PolicyDetailsProperty{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			CrossRegionCopy: []interface{}{
//   				&CrossRegionCopyActionProperty{
//   					EncryptionConfiguration: &EncryptionConfigurationProperty{
//   						CmkArn: jsii.String("cmkArn"),
//   						Encrypted: jsii.Boolean(false),
//   					},
//   					RetainRule: &CrossRegionCopyRetainRuleProperty{
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   					Target: jsii.String("target"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	CopyTags: jsii.Boolean(false),
//   	CreateInterval: jsii.Number(123),
//   	CrossRegionCopyTargets: crossRegionCopyTargets,
//   	EventSource: &EventSourceProperty{
//   		Parameters: &EventParametersProperty{
//   			DescriptionRegex: jsii.String("descriptionRegex"),
//   			EventType: jsii.String("eventType"),
//   			SnapshotOwner: []*string{
//   				jsii.String("snapshotOwner"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Exclusions: &ExclusionsProperty{
//   		ExcludeBootVolumes: jsii.Boolean(false),
//   		ExcludeTags: excludeTags,
//   		ExcludeVolumeTypes: excludeVolumeTypes,
//   	},
//   	ExtendDeletion: jsii.Boolean(false),
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
//   	PolicyLanguage: jsii.String("policyLanguage"),
//   	PolicyType: jsii.String("policyType"),
//   	ResourceLocations: []*string{
//   		jsii.String("resourceLocations"),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	RetainInterval: jsii.Number(123),
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
//   					CmkArn: jsii.String("cmkArn"),
//   					CopyTags: jsii.Boolean(false),
//   					DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   						Interval: jsii.Number(123),
//   						IntervalUnit: jsii.String("intervalUnit"),
//   					},
//   					Encrypted: jsii.Boolean(false),
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
type CfnLifecyclePolicyPropsMixin_PolicyDetailsProperty struct {
	// *[Event-based policies only]* The actions to be performed when the event-based policy is activated.
	//
	// You can specify only one action per policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-actions
	//
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// *[Default policies only]* Indicates whether the policy should copy tags from the source resource to the snapshot or AMI.
	//
	// If you do not specify a value, the default is `false` .
	//
	// Default: false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-copytags
	//
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// *[Default policies only]* Specifies how often the policy should run and create snapshots or AMIs.
	//
	// The creation frequency can range from 1 to 7 days. If you do not specify a value, the default is 1.
	//
	// Default: 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-createinterval
	//
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// *[Default policies only]* Specifies destination Regions for snapshot or AMI copies.
	//
	// You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-crossregioncopytargets
	//
	CrossRegionCopyTargets interface{} `field:"optional" json:"crossRegionCopyTargets" yaml:"crossRegionCopyTargets"`
	// *[Event-based policies only]* The event that activates the event-based policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-eventsource
	//
	EventSource interface{} `field:"optional" json:"eventSource" yaml:"eventSource"`
	// *[Default policies only]* Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.
	//
	// The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-exclusions
	//
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
	// *[Default policies only]* Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.
	//
	// By default ( *ExtendDeletion=false* ):
	//
	// - If a source resource is deleted, Amazon Data Lifecycle Manager will continue to delete previously created snapshots or AMIs, up to but not including the last one, based on the specified retention period. If you want Amazon Data Lifecycle Manager to delete all snapshots or AMIs, including the last one, specify `true` .
	// - If a policy enters the error, disabled, or deleted state, Amazon Data Lifecycle Manager stops deleting snapshots and AMIs. If you want Amazon Data Lifecycle Manager to continue deleting snapshots or AMIs, including the last one, if the policy enters one of these states, specify `true` .
	//
	// If you enable extended deletion ( *ExtendDeletion=true* ), you override both default behaviors simultaneously.
	//
	// If you do not specify a value, the default is `false` .
	//
	// Default: false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-extenddeletion
	//
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// *[Custom snapshot and AMI policies only]* A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// > If you are modifying a policy that was created or previously modified using the Amazon Data Lifecycle Manager console, then you must include this parameter and specify either the default values or the new values that you require. You can't omit this parameter or set its values to null.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The type of policy to create. Specify one of the following:.
	//
	// - `SIMPLIFIED` To create a default policy.
	// - `STANDARD` To create a custom policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-policylanguage
	//
	PolicyLanguage *string `field:"optional" json:"policyLanguage" yaml:"policyLanguage"`
	// The type of policy.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account .
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// *[Custom snapshot and AMI policies only]* The location of the resources to backup.
	//
	// - If the source resources are located in a Region, specify `CLOUD` . In this case, the policy targets all resources of the specified type with matching target tags across all Availability Zones in the Region.
	// - *[Custom snapshot policies only]* If the source resources are located in a Local Zone, specify `LOCAL_ZONE` . In this case, the policy targets all resources of the specified type with matching target tags across all Local Zones in the Region.
	// - If the source resources are located on an Outpost in your account, specify `OUTPOST` . In this case, the policy targets all resources of the specified type with matching target tags across all of the Outposts in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcelocations
	//
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// *[Default policies only]* Specify the type of default policy to create.
	//
	// - To create a default policy for EBS snapshots, that creates snapshots of all volumes in the Region that do not have recent backups, specify `VOLUME` .
	// - To create a default policy for EBS-backed AMIs, that creates EBS-backed AMIs from all instances in the Region that do not have recent backups, specify `INSTANCE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// *[Custom snapshot policies only]* The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// *[Default policies only]* Specifies how long the policy should retain snapshots or AMIs before deleting them.
	//
	// The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time. If you do not specify a value, the default is 7.
	//
	// Default: 7.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-policydetails.html#cfn-dlm-lifecyclepolicy-policydetails-retaininterval
	//
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
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

