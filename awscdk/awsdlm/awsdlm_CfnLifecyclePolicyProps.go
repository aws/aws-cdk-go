package awsdlm

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLifecyclePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecyclePolicyProps := &CfnLifecyclePolicyProps{
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
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
//   		PolicyType: jsii.String("policyType"),
//   		ResourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
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
//   	State: jsii.String("state"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicyProps struct {
	// A description of the lifecycle policy.
	//
	// The characters ^[0-9A-Za-z _-]+$ are supported.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The configuration details of the lifecycle policy.
	PolicyDetails interface{} `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// The activation state of the lifecycle policy.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags to apply to the lifecycle policy during creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

