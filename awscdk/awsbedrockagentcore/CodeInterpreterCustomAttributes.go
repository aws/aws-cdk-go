package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes for specifying an imported Code Interpreter Custom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//
//   codeInterpreterCustomAttributes := &CodeInterpreterCustomAttributes{
//   	CodeInterpreterArn: jsii.String("codeInterpreterArn"),
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
type CodeInterpreterCustomAttributes struct {
	// The ARN of the agent.
	CodeInterpreterArn *string `field:"required" json:"codeInterpreterArn" yaml:"codeInterpreterArn"`
	// The ARN of the IAM role associated to the code interpreter.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The created timestamp of the code interpreter.
	// Default: undefined - No created timestamp is provided.
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// When this code interpreter was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	LastUpdatedAt *string `field:"optional" json:"lastUpdatedAt" yaml:"lastUpdatedAt"`
	// The security groups for this code interpreter, if in a VPC.
	// Default: - By default, the code interpreter is not in a VPC.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The status of the code interpreter.
	// Default: undefined - No status is provided.
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

