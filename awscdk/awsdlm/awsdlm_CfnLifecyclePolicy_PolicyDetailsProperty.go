package awsdlm


// Specifies the configuration of a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDetailsProperty := &policyDetailsProperty{
//   	actions: []interface{}{
//   		&actionProperty{
//   			crossRegionCopy: []interface{}{
//   				&crossRegionCopyActionProperty{
//   					encryptionConfiguration: &encryptionConfigurationProperty{
//   						encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						cmkArn: jsii.String("cmkArn"),
//   					},
//   					target: jsii.String("target"),
//
//   					// the properties below are optional
//   					retainRule: &crossRegionCopyRetainRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   				},
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   	eventSource: &eventSourceProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		parameters: &eventParametersProperty{
//   			eventType: jsii.String("eventType"),
//   			snapshotOwner: []*string{
//   				jsii.String("snapshotOwner"),
//   			},
//
//   			// the properties below are optional
//   			descriptionRegex: jsii.String("descriptionRegex"),
//   		},
//   	},
//   	parameters: &parametersProperty{
//   		excludeBootVolume: jsii.Boolean(false),
//   		excludeDataVolumeTags: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		noReboot: jsii.Boolean(false),
//   	},
//   	policyType: jsii.String("policyType"),
//   	resourceLocations: []*string{
//   		jsii.String("resourceLocations"),
//   	},
//   	resourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	schedules: []interface{}{
//   		&scheduleProperty{
//   			archiveRule: &archiveRuleProperty{
//   				retainRule: &archiveRetainRuleProperty{
//   					retentionArchiveTier: &retentionArchiveTierProperty{
//   						count: jsii.Number(123),
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   				},
//   			},
//   			copyTags: jsii.Boolean(false),
//   			createRule: &createRuleProperty{
//   				cronExpression: jsii.String("cronExpression"),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   				location: jsii.String("location"),
//   				times: []*string{
//   					jsii.String("times"),
//   				},
//   			},
//   			crossRegionCopyRules: []interface{}{
//   				&crossRegionCopyRuleProperty{
//   					encrypted: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					cmkArn: jsii.String("cmkArn"),
//   					copyTags: jsii.Boolean(false),
//   					deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   					retainRule: &crossRegionCopyRetainRuleProperty{
//   						interval: jsii.Number(123),
//   						intervalUnit: jsii.String("intervalUnit"),
//   					},
//   					target: jsii.String("target"),
//   					targetRegion: jsii.String("targetRegion"),
//   				},
//   			},
//   			deprecateRule: &deprecateRuleProperty{
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			fastRestoreRule: &fastRestoreRuleProperty{
//   				availabilityZones: []*string{
//   					jsii.String("availabilityZones"),
//   				},
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			name: jsii.String("name"),
//   			retainRule: &retainRuleProperty{
//   				count: jsii.Number(123),
//   				interval: jsii.Number(123),
//   				intervalUnit: jsii.String("intervalUnit"),
//   			},
//   			shareRules: []interface{}{
//   				&shareRuleProperty{
//   					targetAccounts: []*string{
//   						jsii.String("targetAccounts"),
//   					},
//   					unshareInterval: jsii.Number(123),
//   					unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   				},
//   			},
//   			tagsToAdd: []interface{}{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			variableTags: []interface{}{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	targetTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicy_PolicyDetailsProperty struct {
	// The actions to be performed when the event-based policy is triggered. You can specify only one action per policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The event that triggers the event-based policy.
	//
	// This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter.
	EventSource interface{} `field:"optional" json:"eventSource" yaml:"eventSource"`
	// A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	//
	// > If you are modifying a policy that was created or previously modified using the Amazon Data Lifecycle Manager console, then you must include this parameter and specify either the default values or the new values that you require. You can't omit this parameter or set its values to null.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The valid target resource types and actions a policy can manage.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account .
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT` .
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// The location of the resources to backup.
	//
	// If the source resources are located in an AWS Region , specify `CLOUD` . If the source resources are located on an Outpost in your account, specify `OUTPOST` .
	//
	// If you specify `OUTPOST` , Amazon Data Lifecycle Manager backs up all resources of the specified type with matching target tags across all of the Outposts in your account.
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	//
	// A policy can have up to four schedules—one mandatory schedule and up to three optional schedules.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	Schedules interface{} `field:"optional" json:"schedules" yaml:"schedules"`
	// The single tag that identifies targeted resources for this policy.
	//
	// This parameter is required for snapshot and AMI policies only. If you are creating an event-based policy, omit this parameter.
	TargetTags interface{} `field:"optional" json:"targetTags" yaml:"targetTags"`
}

