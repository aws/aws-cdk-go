package awswellarchitected

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAgentProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentProfileProps := &CfnAgentProfileProps{
//   	AggregationConfiguration: []interface{}{
//   		&AggregationConfigurationProperty{
//   			AccessRoleArn: jsii.String("accessRoleArn"),
//   			AccountId: jsii.String("accountId"),
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	Pillars: []*string{
//   		jsii.String("pillars"),
//   	},
//
//   	// the properties below are optional
//   	BusinessOverview: jsii.String("businessOverview"),
//   	DeletionProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html
//
type CfnAgentProfileProps struct {
	// The aggregation configuration entries (account, regions, access role) associated with this profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-aggregationconfiguration
	//
	AggregationConfiguration interface{} `field:"required" json:"aggregationConfiguration" yaml:"aggregationConfiguration"`
	// The ARN of the IAM role assumed to execute recommendation actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the profile.
	//
	// Unique within the account and used as the last component of the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Well-Architected pillars to focus on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-pillars
	//
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// A business overview for the profile used to improve recommendation quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-businessoverview
	//
	BusinessOverview *string `field:"optional" json:"businessOverview" yaml:"businessOverview"`
	// Whether deletion protection is enabled for the profile.
	//
	// When enabled, the profile cannot be deleted until deletion protection is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Key-value pairs to associate with the Agent Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentprofile.html#cfn-wellarchitected-agentprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

