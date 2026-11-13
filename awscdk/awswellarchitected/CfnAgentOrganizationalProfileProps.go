package awswellarchitected

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAgentOrganizationalProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentOrganizationalProfileProps := &CfnAgentOrganizationalProfileProps{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html
//
type CfnAgentOrganizationalProfileProps struct {
	// The name of the organizational profile.
	//
	// Unique within the account and used as the last component of the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Well-Architected pillars to focus on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-pillars
	//
	Pillars *[]*string `field:"required" json:"pillars" yaml:"pillars"`
	// A business overview for the organizational profile used to improve recommendation quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-businessoverview
	//
	BusinessOverview *string `field:"optional" json:"businessOverview" yaml:"businessOverview"`
	// Whether deletion protection is enabled for the organizational profile.
	//
	// When enabled, the profile cannot be deleted until deletion protection is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of the organizational profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The human-readable display name of the organizational profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Key-value pairs to associate with the Agent Organizational Profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-agentorganizationalprofile.html#cfn-wellarchitected-agentorganizationalprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

