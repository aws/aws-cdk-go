package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes for specifying an imported Browser Custom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//
//   browserCustomAttributes := &BrowserCustomAttributes{
//   	BrowserArn: jsii.String("browserArn"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	LastUpdatedAt: jsii.String("lastUpdatedAt"),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Status: jsii.String("status"),
//   }
//
// Experimental.
type BrowserCustomAttributes struct {
	// The ARN of the agent.
	// Experimental.
	BrowserArn *string `field:"required" json:"browserArn" yaml:"browserArn"`
	// The ARN of the IAM role associated to the browser.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The created timestamp of the browser.
	// Default: undefined - No created timestamp is provided.
	//
	// Experimental.
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// When this browser was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	// Experimental.
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The security groups for this browser, if in a VPC.
	// Default: - By default, the browser is not in a VPC.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The status of the browser.
	// Default: undefined - No status is provided.
	//
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

