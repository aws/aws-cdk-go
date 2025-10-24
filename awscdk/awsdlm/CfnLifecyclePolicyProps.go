package awsdlm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLifecyclePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var crossRegionCopyTargets interface{}
//   var excludeTags interface{}
//   var excludeVolumeTypes interface{}
//
//   cfnLifecyclePolicyProps := &CfnLifecyclePolicyProps{
//   	CopyTags: jsii.Boolean(false),
//   	CreateInterval: jsii.Number(123),
//   	CrossRegionCopyTargets: crossRegionCopyTargets,
//   	DefaultPolicy: jsii.String("defaultPolicy"),
//   	Description: jsii.String("description"),
//   	Exclusions: &ExclusionsProperty{
//   		ExcludeBootVolumes: jsii.Boolean(false),
//   		ExcludeTags: excludeTags,
//   		ExcludeVolumeTypes: excludeVolumeTypes,
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ExtendDeletion: jsii.Boolean(false),
//   	PolicyDetails: &PolicyDetailsProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				CrossRegionCopy: []interface{}{
//   					&CrossRegionCopyActionProperty{
//   						EncryptionConfiguration: &EncryptionConfigurationProperty{
//   							Encrypted: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							CmkArn: jsii.String("cmkArn"),
//   						},
//   						Target: jsii.String("target"),
//
//   						// the properties below are optional
//   						RetainRule: &CrossRegionCopyRetainRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		CopyTags: jsii.Boolean(false),
//   		CreateInterval: jsii.Number(123),
//   		CrossRegionCopyTargets: crossRegionCopyTargets,
//   		EventSource: &EventSourceProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Parameters: &EventParametersProperty{
//   				EventType: jsii.String("eventType"),
//   				SnapshotOwner: []*string{
//   					jsii.String("snapshotOwner"),
//   				},
//
//   				// the properties below are optional
//   				DescriptionRegex: jsii.String("descriptionRegex"),
//   			},
//   		},
//   		Exclusions: &ExclusionsProperty{
//   			ExcludeBootVolumes: jsii.Boolean(false),
//   			ExcludeTags: excludeTags,
//   			ExcludeVolumeTypes: excludeVolumeTypes,
//   		},
//   		ExtendDeletion: jsii.Boolean(false),
//   		Parameters: &ParametersProperty{
//   			ExcludeBootVolume: jsii.Boolean(false),
//   			ExcludeDataVolumeTags: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			NoReboot: jsii.Boolean(false),
//   		},
//   		PolicyLanguage: jsii.String("policyLanguage"),
//   		PolicyType: jsii.String("policyType"),
//   		ResourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		RetainInterval: jsii.Number(123),
//   		Schedules: []interface{}{
//   			&ScheduleProperty{
//   				ArchiveRule: &ArchiveRuleProperty{
//   					RetainRule: &ArchiveRetainRuleProperty{
//   						RetentionArchiveTier: &RetentionArchiveTierProperty{
//   							Count: jsii.Number(123),
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				CopyTags: jsii.Boolean(false),
//   				CreateRule: &CreateRuleProperty{
//   					CronExpression: jsii.String("cronExpression"),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   					Location: jsii.String("location"),
//   					Scripts: []interface{}{
//   						&ScriptProperty{
//   							ExecuteOperationOnScriptFailure: jsii.Boolean(false),
//   							ExecutionHandler: jsii.String("executionHandler"),
//   							ExecutionHandlerService: jsii.String("executionHandlerService"),
//   							ExecutionTimeout: jsii.Number(123),
//   							MaximumRetryCount: jsii.Number(123),
//   							Stages: []*string{
//   								jsii.String("stages"),
//   							},
//   						},
//   					},
//   					Times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				CrossRegionCopyRules: []interface{}{
//   					&CrossRegionCopyRuleProperty{
//   						Encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						CmkArn: jsii.String("cmkArn"),
//   						CopyTags: jsii.Boolean(false),
//   						DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   						RetainRule: &CrossRegionCopyRetainRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   						Target: jsii.String("target"),
//   						TargetRegion: jsii.String("targetRegion"),
//   					},
//   				},
//   				DeprecateRule: &DeprecateRuleProperty{
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				FastRestoreRule: &FastRestoreRuleProperty{
//   					AvailabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				Name: jsii.String("name"),
//   				RetainRule: &RetainRuleProperty{
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				ShareRules: []interface{}{
//   					&ShareRuleProperty{
//   						TargetAccounts: []*string{
//   							jsii.String("targetAccounts"),
//   						},
//   						UnshareInterval: jsii.Number(123),
//   						UnshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   					},
//   				},
//   				TagsToAdd: []interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				VariableTags: []interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		TargetTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RetainInterval: jsii.Number(123),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html
//
type CfnLifecyclePolicyProps struct {
	// *[Default policies only]* Indicates whether the policy should copy tags from the source resource to the snapshot or AMI.
	//
	// If you do not specify a value, the default is `false` .
	//
	// Default: false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-copytags
	//
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// *[Default policies only]* Specifies how often the policy should run and create snapshots or AMIs.
	//
	// The creation frequency can range from 1 to 7 days. If you do not specify a value, the default is 1.
	//
	// Default: 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-createinterval
	//
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// *[Default policies only]* Specifies destination Regions for snapshot or AMI copies.
	//
	// You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-crossregioncopytargets
	//
	CrossRegionCopyTargets interface{} `field:"optional" json:"crossRegionCopyTargets" yaml:"crossRegionCopyTargets"`
	// *[Default policies only]* Specify the type of default policy to create.
	//
	// - To create a default policy for EBS snapshots, that creates snapshots of all volumes in the Region that do not have recent backups, specify `VOLUME` .
	// - To create a default policy for EBS-backed AMIs, that creates EBS-backed AMIs from all instances in the Region that do not have recent backups, specify `INSTANCE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-defaultpolicy
	//
	DefaultPolicy *string `field:"optional" json:"defaultPolicy" yaml:"defaultPolicy"`
	// A description of the lifecycle policy.
	//
	// The characters ^[0-9A-Za-z _-]+$ are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// *[Default policies only]* Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.
	//
	// The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-exclusions
	//
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-extenddeletion
	//
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// The configuration details of the lifecycle policy.
	//
	// > If you create a default policy, you can specify the request parameters either in the request body, or in the PolicyDetails request structure, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-policydetails
	//
	PolicyDetails interface{} `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// *[Default policies only]* Specifies how long the policy should retain snapshots or AMIs before deleting them.
	//
	// The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time. If you do not specify a value, the default is 7.
	//
	// Default: 7.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-retaininterval
	//
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
	// The activation state of the lifecycle policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags to apply to the lifecycle policy during creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html#cfn-dlm-lifecyclepolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

