package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes for importing an existing Agent Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//
//   agentRuntimeAttributes := &AgentRuntimeAttributes{
//   	AgentRuntimeArn: jsii.String("agentRuntimeArn"),
//   	AgentRuntimeId: jsii.String("agentRuntimeId"),
//   	AgentRuntimeName: jsii.String("agentRuntimeName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AgentRuntimeVersion: jsii.String("agentRuntimeVersion"),
//   	AgentStatus: jsii.String("agentStatus"),
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type AgentRuntimeAttributes struct {
	// The ARN of the agent runtime.
	// Experimental.
	AgentRuntimeArn *string `field:"required" json:"agentRuntimeArn" yaml:"agentRuntimeArn"`
	// The ID of the agent runtime.
	// Experimental.
	AgentRuntimeId *string `field:"required" json:"agentRuntimeId" yaml:"agentRuntimeId"`
	// The name of the agent runtime.
	// Experimental.
	AgentRuntimeName *string `field:"required" json:"agentRuntimeName" yaml:"agentRuntimeName"`
	// The IAM role ARN.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The version of the agent runtime When importing a runtime and this is not specified or undefined, endpoints created on this runtime will point to version "1" unless explicitly overridden.
	// Default: - undefined.
	//
	// Experimental.
	AgentRuntimeVersion *string `field:"optional" json:"agentRuntimeVersion" yaml:"agentRuntimeVersion"`
	// The current status of the agent runtime.
	// Default: - Status not available.
	//
	// Experimental.
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// The time at which the runtime was created.
	// Default: - Creation time not available.
	//
	// Experimental.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The description of the agent runtime.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time at which the runtime was last updated.
	// Default: - Last update time not available.
	//
	// Experimental.
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The security groups for this runtime, if in a VPC.
	// Default: - By default, the runtime is not in a VPC.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

