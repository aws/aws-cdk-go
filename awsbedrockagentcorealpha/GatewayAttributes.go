package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Attributes for importing an existing Gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//
//   gatewayAttributes := &GatewayAttributes{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	GatewayId: jsii.String("gatewayId"),
//   	GatewayName: jsii.String("gatewayName"),
//   	Role: role,
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayAttributes struct {
	// The ARN of the gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The ID of the gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The name of the gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The IAM role that provides permissions for the gateway.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

