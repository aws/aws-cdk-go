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
//   cfnLifecyclePolicyProps := &cfnLifecyclePolicyProps{
//   	description: jsii.String("description"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	policyDetails: &policyDetailsProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				crossRegionCopy: []interface{}{
//   					&crossRegionCopyActionProperty{
//   						encryptionConfiguration: &encryptionConfigurationProperty{
//   							encrypted: jsii.Boolean(false),
//
//   							// the properties below are optional
//   							cmkArn: jsii.String("cmkArn"),
//   						},
//   						target: jsii.String("target"),
//
//   						// the properties below are optional
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   		eventSource: &eventSourceProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			parameters: &eventParametersProperty{
//   				eventType: jsii.String("eventType"),
//   				snapshotOwner: []*string{
//   					jsii.String("snapshotOwner"),
//   				},
//
//   				// the properties below are optional
//   				descriptionRegex: jsii.String("descriptionRegex"),
//   			},
//   		},
//   		parameters: &parametersProperty{
//   			excludeBootVolume: jsii.Boolean(false),
//   			excludeDataVolumeTags: []interface{}{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			noReboot: jsii.Boolean(false),
//   		},
//   		policyType: jsii.String("policyType"),
//   		resourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		schedules: []interface{}{
//   			&scheduleProperty{
//   				archiveRule: &archiveRuleProperty{
//   					retainRule: &archiveRetainRuleProperty{
//   						retentionArchiveTier: &retentionArchiveTierProperty{
//   							count: jsii.Number(123),
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				copyTags: jsii.Boolean(false),
//   				createRule: &createRuleProperty{
//   					cronExpression: jsii.String("cronExpression"),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   					location: jsii.String("location"),
//   					times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				crossRegionCopyRules: []interface{}{
//   					&crossRegionCopyRuleProperty{
//   						encrypted: jsii.Boolean(false),
//
//   						// the properties below are optional
//   						cmkArn: jsii.String("cmkArn"),
//   						copyTags: jsii.Boolean(false),
//   						deprecateRule: &crossRegionCopyDeprecateRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						retainRule: &crossRegionCopyRetainRuleProperty{
//   							interval: jsii.Number(123),
//   							intervalUnit: jsii.String("intervalUnit"),
//   						},
//   						target: jsii.String("target"),
//   						targetRegion: jsii.String("targetRegion"),
//   					},
//   				},
//   				deprecateRule: &deprecateRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				fastRestoreRule: &fastRestoreRuleProperty{
//   					availabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				name: jsii.String("name"),
//   				retainRule: &retainRuleProperty{
//   					count: jsii.Number(123),
//   					interval: jsii.Number(123),
//   					intervalUnit: jsii.String("intervalUnit"),
//   				},
//   				shareRules: []interface{}{
//   					&shareRuleProperty{
//   						targetAccounts: []*string{
//   							jsii.String("targetAccounts"),
//   						},
//   						unshareInterval: jsii.Number(123),
//   						unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   					},
//   				},
//   				tagsToAdd: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				variableTags: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		targetTags: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	state: jsii.String("state"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

